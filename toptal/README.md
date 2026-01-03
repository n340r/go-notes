# Book Shop API

This is a simple bookshop REST API that allows users to browse books, put them in the cart, and purchase them (without actual payment processing).

## Task scope and expectations

There are 3 permission levels:

- Anonymous users: they should be able to list and filter books by category. It should also be possible to filter by multiple categories, which would return books from all of the categories selected.
- Authenticated users: they should be able to do everything anonymous users can, but can also add books to the cart. After that, they should be able to check out and buy books.
- Administrators: they should be able to CRUD categories and books.

## Functional requirements
- It should be possible for users to register and authenticate using email+password through the API.
- Users can be made admins only through the DB.
- Admins can CRUD categories. Every category has a name and books assigned to it. Categories hierarchy is flat - meaning that they can’t be nested.
- Admins can CRUD books. Every book has a title, year published, author name, price in USD, and category. Each book can (and must) be assigned to a category. Every book also has a number of copies in stock. Books that are sold out should not be visible in the listing, and it should not be possible to buy them. Stock can be specified when a book is created, but can’t be edited later.
- Visitors (including unauthenticated ones) should be able to browse and filter books.
- Authenticated users should be able to add books to their cart. For simplicity, let’s assume that users can buy multiple books, but only one copy of each (so quantity is not necessary).
- There should be an endpoint that completes checkout and “buys” books currently in the cart. Please note that for simplicity, this endpoint should not take in any credit card details. It should simply pretend that it received them and can assume that a payment was made successfully, and should simply clear the cart and reduce the available quantity of books bought.
- Be mindful of race condition cases where two users might want to buy the same book when there is only one in stock. This should not be possible.
- Note that if a user adds a book to the cart and doesn’t buy it, after 30 minutes it needs to be “released” from the user's cart and made available again for the others.

## Technical requirements
- The server needs to expose a GraphQL or RESTful API. It needs to be able to support non-web clients such as mobile apps.
- Be mindful of the edge cases and unexpected scenarios.
- Be mindful of security and data validation.
- It is expected that this system can handle 10 000+ books being offered.

## How to run
- `make dc` runs docker-compose with the app container on port 8080 for you.
- `make test` runs the tests
- `make run` runs the app locally on port 8080 without docker.
- `make lint` runs the linter

Short project demo: https://www.dropbox.com/s/g9592xfj5wxbpmo/toptal_book_shop_api_demo.mp4?dl=0

## Solution notes
- :trident: clean architecture (handler->service->repository)
- :book: standard Go project layout (well, more or less :blush:)
- :cd: docker compose + Makefile included
- :card_file_box: PostgreSQL migrations included
- :heavy_check_mark: Postman collection included
- :lock: race conditions are handled by transactions and `SELECT ... FOR UPDATE` in the SQL queries