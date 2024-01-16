package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name                    string                  `json:"name"`
	FeaturesImg             string                  `json:"featuresImg"`
	FeatureOne              string                  `json:"featureOne"`
	FeatureTwo              string                  `json:"featureTwo"`
	FeatureThree            string                  `json:"featureThree"`
	ProductContainerContent ProductContainerContent `json:"productContainerContent"`
	ProductCarouselItems    []ProductCarouselItem   `json:"productCarouselItems"`
	SubProducts             []SubProduct            `json:"subProducts"`
}

type SubProduct struct {
	gorm.Model
	ProductID uint   `json:"productId"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Image     string `json:"imageUrl"`
}

type ProductCarouselItem struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Image     string `json:"image"`
}

type ProductContainerContent struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Video     string `json:"videUrl"`
}

func (p *Product) SaveToDB() (*Product, error) {
	err := db.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

// findAll products
func (p *Product) GetAllProducts() ([]Product, error) {
	var products []Product
	err := db.Debug().Table("products").Find(&products).Error
	if err != nil {
		return []Product{}, err
	}
	// product for loop
	for i, _ := range products {
		err := db.Debug().Table("product_container_contents").Where("product_id=?", products[i].ID).Find(&products[i].ProductContainerContent).Error
		if err != nil {
			return []Product{}, err
		}

		err = db.Debug().Table("product_carousel_items").Where("product_id=?", products[i].ID).Find(&products[i].ProductCarouselItems).Error
		if err != nil {
			return []Product{}, err
		}
		err = db.Debug().Table("sub_products").Where("product_id=?", products[i].ID).Find(&products[i].SubProducts).Error
		if err != nil {
			return []Product{}, err
		}

	}
	return products, nil
}

func (product *Product) GetProductByID(id uint) (*Product, error) {
	err := db.Debug().Table("products").Where("id = ?", id).Find(&product).Error
	if err != nil {
		return &Product{}, err
	}
	if product.ID == uint(id) {
		err := db.Debug().Table("product_container_contents").Where("product_id=?", id).Find(&product.ProductContainerContent).Error
		if err != nil {
			return &Product{}, err
		}

		err = db.Debug().Table("product_carousel_items").Where("product_id=?", id).Find(&product.ProductCarouselItems).Error
		if err != nil {
			return &Product{}, err
		}
		err = db.Debug().Table("sub_products").Where("product_id=?", id).Find(&product.SubProducts).Error
		if err != nil {
			return &Product{}, err
		}

	}
	return product, nil
}

// update product
func (p *Product) UpdateProduct(pid uint64) (*Product, error) {
	err := db.Debug().Table("products").Where("id = ?", pid).Updates(Product{Name: p.Name, FeaturesImg: p.FeaturesImg, FeatureOne: p.FeatureOne, FeatureTwo: p.FeatureTwo, FeatureThree: p.FeatureThree}).Error
	if err != nil {
		return &Product{}, err
	}
	for _, productCarouselItem := range p.ProductCarouselItems {
		db.Model(&ProductCarouselItem{}).Where("id = ?", productCarouselItem.ID).Updates(map[string]interface{}{
			"title":  productCarouselItem.Title,
			"detail": productCarouselItem.Detail,
			"image":  productCarouselItem.Image,
		})
	}
	db.Model(&ProductContainerContent{}).Where("id = ?", p.ProductContainerContent.ID).Updates(map[string]interface{}{
		"title":  p.ProductContainerContent.Title,
		"detail": p.ProductContainerContent.Detail,
		"video":  p.ProductContainerContent.Video,
	})

	for _, subproduct := range p.SubProducts {
		db.Model(&SubProduct{}).Where("id = ?", subproduct.ID).Updates(map[string]interface{}{
			"title":  subproduct.Title,
			"detail": subproduct.Detail,
			"image":  subproduct.Image,
		})
	}
	return p, nil
}

// delete productbyid
func (p Product) RemoveFromDb(pid uint64) error {
	err := db.Debug().Table("products").Where("id = ?", pid).Delete(Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

// Carousel item
func (pc *ProductCarouselItem) SaveToDb() (*ProductCarouselItem, error) {
	err := db.Create(&pc).Error
	if err != nil {
		return &ProductCarouselItem{}, err
	}
	return pc, nil
}

// get by product id
func GetProductCarouselItemsByProductId(id uint) ([]ProductCarouselItem, error) {
	var productCarouselItems []ProductCarouselItem
	err := db.Where("product_id = ?", id).Find(&productCarouselItems).Error
	if err != nil {
		return []ProductCarouselItem{}, err
	}
	return productCarouselItems, nil
}

// remove from db
func (pc *ProductCarouselItem) RemoveFromDb() (*ProductCarouselItem, error) {
	err := db.Delete(&pc).Error
	if err != nil {
		return &ProductCarouselItem{}, err
	}
	return pc, nil
}

// product container content
func (pc *ProductContainerContent) SaveToDb() (*ProductContainerContent, error) {
	err := db.Create(&pc).Error
	if err != nil {
		return &ProductContainerContent{}, err
	}
	return pc, nil
}

// get by product id
func GetProductContainerContentsByProductId(id uint) ([]ProductContainerContent, error) {
	var productContainerContents []ProductContainerContent
	err := db.Where("product_id = ?", id).Find(&productContainerContents).Error
	if err != nil {
		return []ProductContainerContent{}, err
	}
	return productContainerContents, nil
}

// remove from db
func (pc *ProductContainerContent) RemoveFromDb() (*ProductContainerContent, error) {
	err := db.Delete(&pc).Error
	if err != nil {
		return &ProductContainerContent{}, err
	}
	return pc, nil
}

// home cover
func (sp *SubProduct) SaveToDb() (*SubProduct, error) {
	err := db.Debug().Create(&sp).Error
	if err != nil {
		return &SubProduct{}, err
	}
	return sp, nil
}

// remove from db
func (sp *SubProduct) RemoveFromDb() (*SubProduct, error) {
	err := db.Delete(&sp).Error
	if err != nil {
		return &SubProduct{}, err
	}
	return sp, nil
}

// get All home covers
func GetAllSubProducts() ([]SubProduct, error) {
	var subProducts []SubProduct
	err := db.Find(&subProducts).Error
	if err != nil {
		return []SubProduct{}, nil
	}
	return subProducts, nil
}
