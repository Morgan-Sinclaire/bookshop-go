# bookshop-go

A simple book store API in need of input validation/sanitization.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:

- [Javascript](https://github.com/andey-robins/bookshop-js)
- [Rust](https://github.com/andey-robins/bookshop-rs)

## Versioning

`bookshop-go` is buit with:

- go version go1.19.3 darwin/arm64

## Usage

Start the api using `go run main.go`.

I recommend using [`httpie`](https://httpie.io) for testing of HTTP endpoints on the terminal. Tutorials are available elsewhere online, and you're free to use whatever tools you deem appropriate for testing your code.

## Analysis of Existing Code

First of all, this original code was missing a lot of needed "rows.Next()" lines that caused errors.
However, as written it is safe against SQL injections. For instance, if CreateBook() were written as:

database.Exec(fmt.Sprintf(`INSERT INTO Books (title, author, price) VALUES (%s, %s, %s);`, title, author, price))

then anything could be inputted into the title, author, and price fields, and db.Exec() would simply run it blindly.
However, the code instead parametrizes these fields within that function:

database.Exec(`INSERT INTO Books (title, author, price) VALUES (?, ?, ?);`, title, author, price)

When given multiple parameters, db.Exec() automatically sanitizes the optional parameters for SQL injections, or if they're simply too long.

On the other hand, the price field could be any float32. But we don't want negative prices, so we must manually validate this.

Furthermore, we can make the following post request successfully:

{
  "Title": "Dune",
  "Author": "Frank Herbert"
}

If we do so, it will add that book to the database with a price of 0. This is because we didn't give a price,
and the existing handlers will have it default to 0. This is bad, since we want the listing for each book
to be complete and accurate.

The code we have now validates that there are no missing fields, and that a book's price is positive.

Acknowledgements: I have no idea about JSON stuff. I worked with Long, but he was unfamiliar with how to do this with Go syntax.
So I asked ChatGPT and it wrote all the code for these "create" functions, I checked that it ran, and Long could see that the code made sense.
He then wrote code inspired by this, but since I have no shame I left it as is, and added some logging.
