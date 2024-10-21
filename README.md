# LinkedGram
## How to run the program:
- **Note**: There are four main functions: Create Post, Comment on Post, Search User, and Create Profile. Due to communication limitations between client and server, you can run one function at a time. To run another, restart the program.
- **Route 1**: Register a new account and create a profile. After registration, you'll be redirected to the Create Profile function.
- **Route 2**: Login with an existing account. After login, you’ll be redirected to the main menu to choose between Create Post, Comment on Post, or Search User.
- **Non-main functions**: Logout, view the news feed, upvote or downvote posts. These can be accessed without restarting the program.

### Steps to Run the Program:
1. Run the `AppStarter.go` file to start the server.
2. For multiple clients, run the `Client.go` file for each client instance.

## Key Features:
- **Login and Registration**: Start by entering a username and password. You can either log in or register a new account, which will redirect you to profile creation.
- **Profile Creation**: Choose from two avatar pictures ("boilermaker" or "cute squirrel"). Upon completing the registration, you'll be redirected to the login page with pre-filled credentials.
- **News Feed**: Displays posts and comments. The feed will be empty until you create posts.
- **Create and Comment on Posts**: Users can create posts and comment on them. Due to client-server communication issues, posts and comments may not appear immediately without restarting.
- **Search Users**: Allows searching for users by username to view their profile and avatar.
- **Voting**: Upvote or downvote posts. The functionality is ready for future enhancements.
- **Logout**: Switch accounts without closing the program.

### AppStarter.go
The entry point for starting both the server and client.

- `main()`: Starts the server, then the client GUI. If errors occur, they are printed to the console.

### Client.go
Handles client-side operations:
- `main()`: Connects to the server and displays the login menu.
- Various methods handle login, registration, searching for users, creating posts, commenting, and managing votes.
