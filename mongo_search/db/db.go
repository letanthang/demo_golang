package db

import (
	"context"

	"github.com/letanthang/demo_golang/mongo_search/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchStudent(req types.StudentSearchReq) (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)

	filter := bson.M{}

	if req.ID != 0 {
		filter["id"] = req.ID
	}

	if req.Email != "" {
		filter["email"] = primitive.Regex{Pattern: req.Email, Options: "i"}
	}

	if req.ClassName != "" {
		filter["class_name"] = primitive.Regex{Pattern: req.ClassName, Options: "i"}
	}

	if req.Name != "" {
		filter["$or"] = bson.A{
			bson.M{"first_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
			bson.M{"last_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
			bson.M{"email": primitive.Regex{Pattern: req.Name, Options: "i"}},
		}
	}

	cur, err := Client.Database(dbName).Collection("student").Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cur.All(context.TODO(), &students); err != nil {
		return nil, err
	}

	return &students, nil
}
