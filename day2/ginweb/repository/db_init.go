package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var topicIndexMap map[int64]*Topic
var postIndexMap map[int64][]*Post

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		err := json.Unmarshal([]byte(text), &topic)
		if err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		err := json.Unmarshal([]byte(text), &post)
		if err != nil {
			return err
		}
		postTmpMap[post.ParentId] = append(postTmpMap[post.ParentId], &post)
	}
	postIndexMap = postTmpMap
	return nil
}
