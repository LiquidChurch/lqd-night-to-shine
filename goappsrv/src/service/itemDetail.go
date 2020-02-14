package service

import (
  "time"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

func PostItemDetail(c helper.ContextDetail, postItemDetail *model.ItemDetail) (*model.ItemDetail, error) {
  if postItemDetail.ID == "" {
    var err error
    postItemDetail.ID, err = helper.GenerateULID()
    postItemDetail.CreatedTime = time.Now()
    if err != nil {
      helper.Log(c, "error", "generate ULID", "error", err.Error())
      return model.NullItemDetail, err
    }
  }
  
  helper.Log(c, "info", "POST Entered", "itemID", postItemDetail.ID)
  
  postItemDetail.UpdatedTime = time.Now()
  
  saveErr := model.SaveItemDetail(c.Ctx, postItemDetail.ID, postItemDetail)    
  if saveErr != nil {
    helper.Log(c, "error", "save item detail", "itemID", postItemDetail.ID, "error", saveErr.Error())
    return model.NullItemDetail, saveErr
  }

  helper.Log(c, "info", "POST Exited", "itemID", postItemDetail.ID)
  return postItemDetail, nil
}
