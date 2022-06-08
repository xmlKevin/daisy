package task

/*
  @Author : xmlKevin
*/

import (
	"context"
	"daisy/pkg/task/worker"
)

func Send(classify string, scriptPath string, params string) {
	worker.SendTask(context.Background(), classify, scriptPath, params)
}
