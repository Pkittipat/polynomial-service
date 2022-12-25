# polynomial-service
## Run Server
#### Setup dependencies
```
$ make setup
```
#### start the server
```
$ make start
```
## API Endpoints
#### POST /polynomial/guess
Take a JSON request body with the folling parameters:
- `x:` the value of first index that missing from the dataset.
- `y:` the value of second index that missing from the dataset.
- `z:` the value of third index that missing from the dataset.

Returns a JSON response with the folling parameters:
- `is_polynomial:` fact of guess x, y, z are correct or not.

#### GET /polynomial/dataset
Returns the dataset of [1, x, 8, 17, y, z, 78, 113] with random x, y and z already.
