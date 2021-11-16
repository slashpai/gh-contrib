# gh-contrib

To count github contributions org wise or repo wise in a org

## Pre-requistes

- Have a [gihub personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) which has permission to read repo status and user info. Takecare not to share token with anyone else since it works same as password :)

  ![personal token](img/personal_token.png)

## Note

Since github api is [rate limited](https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting) application might take longer to execute. I am yet to handle the rate limiting inside code to keep it waiting.

## Features

- [x] Count contributions org wise
- [x] Count contributions repository wise in a org
- [x] Generates file output to share

## TODO

- [ ] Add more logging info
- [ ] General Improvement & Optimisation
- [ ] Add parallelism to code
- [ ] Option to count contributions without specifying org
- [ ] Add web output
- [ ] Create executable to easily install the application
- [ ] Handle rate-limiting which will be an issue for large org's

## Contributing

- Bug reports and pull requests are welcome

  - Fork the project on GitHub
  - Clone the project
  - Add changes (and tests if applicable)
  - Commit and push
  - Create a pull request

## License

[MIT](LICENSE)
