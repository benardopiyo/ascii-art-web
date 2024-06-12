# ASCII Art Web

## Description
This project provides a web interface for generating ASCII art using different banner styles.

## Usage

* To run the server, use the following command:

```bash
go run ./server

```

* Open a web browser and go to 

```bash
http://localhost:8080
```

## Implementation details

* The application uses Go's standard library to create a web server and handle HTTP requests. 
* ASCII art generation is implemented using predefined banner styles.

### Endpoints
- `GET /`: Renders the main page with text input and banner selection.
- `POST /ascii-art`: Generates and displays the ASCII art based on the provided text and banner style.


## Notes

1. Ensure you have Go installed on your system.
2. Only standard Go packages are used as per the requirements.
3. Error handling is implemented to ensure appropriate HTTP status codes are returned.

## Author
[Benard Opiyo](https://learn.zone01kisumu.ke/git/beopiyo)