# GitLab Users Exporter

This Go script fetches user data from a GitLab instance and exports it to a CSV file.

## Prerequisites

- Go installed on your machine
- GitLab instance with API access

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/addhe/get_gitlab_user.git
    cd get_gitlab_user
    ```

2. Set environment variables:

    - `GITLAB_TOKEN`: GitLab private token with API access.
    - `GITLAB_URL`: GitLab instance URL.

    Example:

    ```bash
    export GITLAB_TOKEN="<your-gitlab-token>"
    export GITLAB_URL="https://gitlab.example.com"
    ```

3. Build or run the script:

    ```bash
    go build
    ./<executable-name>
    ```

    ```
    go run get_gitlab_user.go
    ```

## Functionality

The script performs the following tasks:

1. **Fetch GitLab Users:**

    The script fetches a list of GitLab users by making API requests to the GitLab instance.

2. **Export to CSV:**

    The fetched user data is exported to a CSV file named `users.csv`. The CSV file contains the following columns:

    - ID
    - Username
    - Name
    - Email
    - State

## Adjusting Configuration

- `perPage` constant in the script determines the number of users fetched per API request. Modify it based on your needs.

## Notes

- Ensure that the provided GitLab token has the necessary permissions to access user information.

- Adjust the `perPage` constant in the script according to your GitLab instance's capabilities and your preferences for pagination.

## License

This project is licensed under the [MIT License](LICENSE).
