package dataloaders

import (
	"context"
	"net/http"
	"time"

	"gorm.io/gorm"

	"example/models"
)

var (
	EmployeeLoaderByTaskKey    = "employeeloaderbytask"
	EmployeeLoaderByCompanyKey = "employeeloaderbycompany"
	taskLoaderKey              = "taskloader"
	companyLoaderKey           = "companyloader"
)

func DataLoaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		emplyeeLoaderByTask := EmployeeLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []int) (employees [][]*models.Employee, errs []error) {
				taskKeyMap := make(map[int][]*models.Employee, len(keys))
				for i := range keys {
					var emps []*models.Employee
					err := db.Model(models.Employee{}).
						Joins("inner join employee_tasks on employee_id = id").
						Where("task_id = ?", keys[i]).Find(&emps).Error
					if err != nil {
						errs = append(errs, err)
						return
					}
					taskKeyMap[keys[i]] = emps
				}
				result := make([][]*models.Employee, len(keys))
				for i, id := range keys {
					result[i] = taskKeyMap[id]
				}
				return result, nil
			},
		}
		taskLoader := TaskLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []int) (tasks [][]*models.Task, errs []error) {
				employeeKeyMap := make(map[int][]*models.Task, len(keys))
				for i := range keys {
					var tsks []*models.Task
					err := db.Model(models.Task{}).
						Joins("inner join employee_tasks on task_id = id").
						Where("employee_id = ?", keys[i]).Find(&tsks).Error
					if err != nil {
						errs = append(errs, err)
						return
					}
					employeeKeyMap[keys[i]] = tsks
				}
				result := make([][]*models.Task, len(keys))
				for i, id := range keys {
					result[i] = employeeKeyMap[id]
				}
				return result, nil
			},
		}
		emplyeeLoaderByCompany := EmployeeLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []int) (employees [][]*models.Employee, errs []error) {
				companyKeyMap := make(map[int][]*models.Employee, len(keys))
				for i := range keys {
					var emps []*models.Employee
					err := db.Model(models.Employee{}).
						Where("company_id = ?", keys[i]).Find(&emps).Error
					if err != nil {
						errs = append(errs, err)
						return
					}
					companyKeyMap[keys[i]] = emps
				}
				result := make([][]*models.Employee, len(keys))
				for i, id := range keys {
					result[i] = companyKeyMap[id]
				}
				return result, nil
			},
		}
		companyLoader := CompanyLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []int) (companies []*models.Company, errs []error) {
				err := db.Model(models.Company{}).Where("id IN ?", keys).Find(&companies).Error
				if err != nil {
					errs = append(errs, err)
					return
				}
				companyMap := make(map[int]*models.Company, len(keys))
				for i := range companies {
					companyMap[companies[i].ID] = companies[i]
				}
				result := make([]*models.Company, len(keys))
				for i, id := range keys {
					result[i] = companyMap[id]
				}
				return result, nil
			},
		}
		ctx := context.WithValue(r.Context(), EmployeeLoaderByTaskKey, &emplyeeLoaderByTask)
		ctx = context.WithValue(ctx, taskLoaderKey, &taskLoader)
		ctx = context.WithValue(ctx, EmployeeLoaderByCompanyKey, &emplyeeLoaderByCompany)
		ctx = context.WithValue(ctx, companyLoaderKey, &companyLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetEmployeeLoader(ctx context.Context, key string) *EmployeeLoader {
	return ctx.Value(key).(*EmployeeLoader)
}

func GetTaskLoader(ctx context.Context) *TaskLoader {
	return ctx.Value(taskLoaderKey).(*TaskLoader)
}

func GetCompanyLoader(ctx context.Context) *CompanyLoader {
	return ctx.Value(companyLoaderKey).(*CompanyLoader)
}
