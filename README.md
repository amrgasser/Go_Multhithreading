### To Run the tests: `go test ./...`

### To Run the tests: `go run .`

### The application utilizes channels, waitgroups and Mutex to handle concurrent acccess. When a withdrawal or deposit request is done, the user (Account) is locked until the transaction is done. All other go routines on the same accounts have to wait for this request to end.

---

### To access the accounts in O(1) time we used an array to store all the users `users`, and a HashMap `userMap` that maps the ID -> index of user in `users`.
