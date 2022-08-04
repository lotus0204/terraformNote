package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Table struct {
	name string
	db   *dynamodb.DynamoDB
}

// Create session.
func (t *Table) Init(name string) {
	db := dynamodb.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-2")})
	t.name = name
	t.db = db
}

// Get one item
func (t *Table) GetItem(hkName string, hkValue string, from interface{}) (*dynamodb.QueryOutput, error) {
	hash := expression.Key(hkName).Equal(expression.Value(hkValue))
	expr, err := expression.NewBuilder().WithKeyCondition(hash).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return nil, err
	}

	query := dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		TableName:                 aws.String(t.name),
	}
	fmt.Println("111")
	fmt.Println(query)

	exkey, err := dynamodbattribute.MarshalMap(from)

	if err != nil {
		fmt.Println("Got error building ExclusiveStartKey:")
		fmt.Println(err.Error())
		return nil, err
	} else {
		query.ExclusiveStartKey = exkey
	}

	result, err := t.db.Query(&query)

	if err != nil {
		fmt.Println("Got error while query:")
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

// Put item into dynamodb table
func (t *Table) PutItem(item interface{}) error {

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling attribute item:")
		fmt.Println(err.Error())
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(t.name),
	}

	_, err = t.db.PutItem(input)
	if err != nil {
		fmt.Println("Got error PutItem:")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (t *Table) ListItem(hkName string, hkValue string, paginated bool, from interface{}) (*dynamodb.QueryOutput, error) {
	hash := expression.Key(hkName).Equal(expression.Value(hkValue))
	expr, err := expression.NewBuilder().WithKeyCondition(hash).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return nil, err
	}

	query := dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		TableName:                 aws.String(t.name),
		Limit:                     aws.Int64(5), // 한번에 다섯개 까지 읽어올 수 있다.
	} // 페이지네이션 테스트를 위해 값을 작게 설정했다.

	if paginated {
		exkey, err := dynamodbattribute.MarshalMap(from)

		if err != nil {
			fmt.Println("Got error building ExclusiveStartKey:")
			fmt.Println(err.Error())
			return nil, err
		} else {
			query.ExclusiveStartKey = exkey
		}
	}

	result, err := t.db.Query(&query)

	if err != nil {
		fmt.Println("Got error while query:")
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

// Update item in dynamodb table
func (t *Table) UpdateItem(key interface{}, expr expression.Expression) (*dynamodb.UpdateItemOutput, error) {

	k, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		fmt.Println("Got error marshalling key item:")
		fmt.Println(err.Error())
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       k,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(t.name),
		ReturnValues:              aws.String("ALL_NEW"),
		UpdateExpression:          expr.Update(),
	}

	result, err := t.db.UpdateItem(input)
	if err != nil {
		fmt.Println("Got error UpdateItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

// Delete item in dynamodb table
func (t *Table) DeleteItem(key interface{}) error { // 에러가 없으면 잘 삭제된 것이다.

	k, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		fmt.Println("Got error marshalling key item:")
		fmt.Println(err.Error())
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       k,
		TableName: aws.String(t.name),
	}

	_, err = t.db.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error DeleteItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}
