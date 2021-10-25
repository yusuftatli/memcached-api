# Yemek sepeti In memory key-value store REST-API

While developing the project, it was avoided to use external libraries.

## Usage
>deployed to Heroku 
```
key set => https://yemek-sepeti-heroku.herokuapp.com/set?url=https://www.yemeksepeti.com/
ket get => https://yemek-sepeti-heroku.herokuapp.com/get?key=url

```

>for local
```
key set => http://localhost:8000/set?url=www.yemeksepeti.com/
ket get => http://localhost:8000/get?key=url

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


## License
[MIT](https://choosealicense.com/licenses/mit/)