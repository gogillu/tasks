# tasks
Tasks

Sample Run 
--------------
govind@Administrators-MacBook-Pro problem1 % go run main.go < input.txt
main starts
Logic

--------------------
Enter item name : Enter item price : Enter item quantity : Enter item type : ==> new Item Added {item1 100 20 raw 0}

Do you want to enter details of any other item (y/n):
--------------------
Enter item name : Enter item price : Enter item quantity : Enter item type : ==> new Item Added {it2 5 1000 imported 0}

Do you want to enter details of any other item (y/n):
--------------------
Enter item name : Enter item price : Enter item quantity : Enter item type : ==> new Item Added {itm3 1000 21 manufactured 0}

Do you want to enter details of any other item (y/n):
--------------------
Enter item name : Enter item price : Enter item quantity : Enter item type : ==> new Item Added {item4 50 2 imported 0}

Do you want to enter details of any other item (y/n):
+-----------+-----------+----------+----------------+---------------------+
| Item Name |     Price | Quantity |           Type |                 Tax |
+-----------+-----------+----------+----------------+---------------------+
|     item1 |       100 |       20 |            raw |          250.000000 |
+-----------+-----------+----------+----------------+---------------------+
|       it2 |         5 |     1000 |       imported |         1025.000000 |
+-----------+-----------+----------+----------------+---------------------+
|      itm3 |      1000 |       21 |   manufactured |         3097.500000 |
+-----------+-----------+----------+----------------+---------------------+
|     item4 |        50 |        2 |       imported |           20.000000 |
+-----------+-----------+----------+----------------+---------------------+
govind@Administrators-MacBook-Pro problem1 % 



------------
govind@Administrators-MacBook-Pro problem1 % go test ./... -race -covermode=atomic -coverprofile=coverage.out
?       problem1        [no test files]
ok      problem1/concrete       0.460s  coverage: 71.2% of statements
ok      problem1/requirement    0.719s  coverage: 0.0% of statements [no tests to run]
govind@Administrators-MacBook-Pro problem1 % 
-------------