package products

import "amazon/lib/types"



func CollcetProducts(Departments []types.Department) {
	for _, Department := range Departments {
		for _, Category := range Department.Categories {
			if len(Category.SubCategories) == 0 {
				CollectAllPagesProducts(Category.Url)
				continue
			}
		}
	}
}


func CollectAllPagesProducts(url string) {

}