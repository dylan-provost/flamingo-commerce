# 14. June 2018

* Price Fields in Cartitems and Carttotals have been changed:
  * Cartitem:
    * Deleted (Dont use anymore): Price / DiscountAmount / PriceInclTax
    * Now Existing: SinglePrice / SinglePriceInclTax / RowTotal / TaxAmount/ RowTotalInclTax / TotalDiscountAmount / ItemRelatedDiscountAmount / NonItemRelatedDiscountAmount / RowTotalWithItemRelatedDiscount / RowTotalWithItemRelatedDiscountInclTax / RowTotalWithDiscountInclTax
    
  * Carttotal:
    * Deleted: DiscountAmount
    * Now Existing: SubTotal / SubTotalInclTax / SubTotalInclTax /SubTotalWithDiscounts / SubTotalWithDiscountsAndTax / TotalDiscountAmount / TotalNonItemRelatedDiscountAmount 

# 17. April 2019

* Cart Item `UniqueID` is removed
  * `Item.ID` is now supposed to be unique
  * The combination `ID` + `DeliveryCode` is no longer required to identify a cart item
  * For non-unique references of certain backend implementations the new field `Item.ExternalReference` can be used
