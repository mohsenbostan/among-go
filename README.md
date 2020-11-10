# Among-Go

This app will update your Twitter profile when you're online in Discord automatically.

actually, this project was for my own laziness to notify my friends on Twitter when I'm in Among Us to come and play with each other. 😉

### Installation

 - clone the repo with this command:
```
git clone https://github.com/mohsenbostan/among-go.git
```

 - run this command to create a copy from .env.example in a new .env file and fill it with proper data:
```
make env
```

 - run the following command to build the Docker Container
```
make dc-build
```

 - run this command to launch the application:
```
make dc-run
```

 - if you want to launch the application in detached mode, then run this command:
```
make dc-run-d
```
Hurray! now you are running the application successfully. 🥳
