#In memory key-value store REST-API

While developing the project, i avoided to use external libraries.

## Usage
>deployed to Heroku 
```
key value set => https://ys-herokuapp.com/set?url=https://www.yemeksepeti.com/
key value get => https://ys-heroku.herokuapp.com/get?key=url
delete all cache value => https://ys-heroku.herokuapp.com/deleteall

```

>for local
```
key set => http://localhost:8000/set?url=www.yemeksepeti.com/
ket get => http://localhost:8000/get?key=url
delete all cache value => http://localhost:8000/deleteall

```

>run, test and build
```
make run
make test
make build

```


>deploy from to heroku server
```
heroku login
heroku create
git push heroku master

```

>docker run
```
docker build
docker run

```


## License
[MIT](https://choosealicense.com/licenses/mit/)
