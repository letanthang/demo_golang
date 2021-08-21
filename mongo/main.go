package main

import (
	"app/models"
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName   = "goclass3008"
	collName = "student"
)

var Client *mongo.Client

func main() {

	Connect()
	student := models.Student{
		FirstName: "Thang",
		LastName:  "Le",
		Email:     "letanthang@gmail.com",
		ClassName: "golang high performace app",
	}
	AddStudent(&student)
	students, err := SearchStudent(models.StudentSearchReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(students)

}

func Connect() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("MONGO_URI", "mongodb://mongo:mongo@localhost:27017")
	uri := viper.GetString("MONGO_URI")
	// uri = "mongodb+srv://mongoadmin:abcd1234@cluster0.cczrl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	fmt.Println("MONGO_URI", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}

func AddStudent(student *models.Student) error {
	start := time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	student.Id = primitive.NewObjectID()
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	_, err := Client.Database(dbName).Collection(collName).InsertOne(ctx, student)
	if err != nil {
		return errors.WithMessage(err, "AddStudent error")
	}
	fmt.Println(time.Since(start))
	return nil
}

func SearchStudent(req models.StudentSearchReq) (*[]models.Student, error) {
	start := time.Now()
	var students []models.Student
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

	cur, err := Client.Database(dbName).Collection(collName).Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cur.All(context.TODO(), &students); err != nil {
		return nil, err
	}

	fmt.Println(time.Since(start))
	return &students, nil
}
