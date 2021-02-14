package card

import (
	"reflect"
	"testing"
)

func TestSumByCategory(t *testing.T) {
	type args struct {
		transactions []Transaction
		userId       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "TestSumByCategory for user id 1",
			args: args{MakeTransactions(1), 1},
			want: map[string]int64{"Eating Places and Restaurants": 1000000, "Grocery Stores, Supermarkets": 1000000},
		},
		{
			name: "TestSumByCategory for user id 2",
			args: args{MakeTransactions(1), 2},
			want: map[string]int64{"Eating Places and Restaurants": 98000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumByCategory(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumByCategoryChannels(t *testing.T) {
	type args struct {
		transactions []Transaction
		userId       int64
		goroutines   int
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "TestSumByCategory for user id 1",
			args: args{MakeTransactions(1), 1, 10},
			want: map[string]int64{"Eating Places and Restaurants": 1000000, "Grocery Stores, Supermarkets": 1000000},
		},
		{
			name: "TestSumByCategory for user id 2",
			args: args{MakeTransactions(1), 2, 10},
			want: map[string]int64{"Eating Places and Restaurants": 98000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumByCategoryChannels(tt.args.transactions, tt.args.userId, tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByCategoryChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumByCategoryMutex(t *testing.T) {
	type args struct {
		transactions []Transaction
		userId       int64
		goroutines   int
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "TestSumByCategory for user id 1",
			args: args{MakeTransactions(1), 1, 10},
			want: map[string]int64{"Eating Places and Restaurants": 1000000, "Grocery Stores, Supermarkets": 1000000},
		},
		{
			name: "TestSumByCategory for user id 2",
			args: args{MakeTransactions(1), 2, 10},
			want: map[string]int64{"Eating Places and Restaurants": 98000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumByCategoryMutex(tt.args.transactions, tt.args.userId, tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByCategoryMutex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumByCategoryMutexWithoutFunc(t *testing.T) {
	type args struct {
		transactions []Transaction
		userId       int64
		goroutines   int
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "TestSumByCategory for user id 1",
			args: args{MakeTransactions(1), 1, 10},
			want: map[string]int64{"Eating Places and Restaurants": 1000000, "Grocery Stores, Supermarkets": 1000000},
		},
		{
			name: "TestSumByCategory for user id 2",
			args: args{MakeTransactions(1), 2, 10},
			want: map[string]int64{"Eating Places and Restaurants": 98000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumByCategoryMutexWithoutFunc(tt.args.transactions, tt.args.userId, tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByCategoryMutexWithoutFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
