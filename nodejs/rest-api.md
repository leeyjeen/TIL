# Create a REST API With Node.JS and MongoDB

## Project setup
```bash
$ node -v
$ npm -v
$ mongo --version
$ sudo npm install -g turbo-cli
$ turbo version
# rest-api 디렉토리에 프로젝트 기본 구조 셋팅
$ turbo new rest-api

$ cd rest-api
# install dependencies as defined in the package.json
$ npm install

# run devserver
$ turbo devserver

- - - - - - - - - - - - - - -

Turbo dev server running on http://localhost:3000
Open this address in a browswer to view your project.
This is a local environment for testing and is NOT actually live on the internet.
To turn off server: CONTROL-C

- - - - - - - - - - - - - - -

# Turbo 계정 Site 생성 후 Site와 로컬 코드 연결
$ turbo login
$ turbo connect
$ turbo deploy
```

## API Route
`app.js`
```js
// import routes
const index = require('./routes/index')
const api = require('./routes/api')

// set routes
app.use('/', index) // return home page
app.use('/api', api) // sample API Routes
```

`routes/api.js`
```js
const express = require('express')
const router = express.Router()

router.get('/test', (req, res) => {
  res.json({
    confirmation: 'success',
    data: 'this is a test endpoint'
  })
})

module.exports = router
```

<img src="./images/02.nodejs_api_get_test.png" width="70%" alt="REST API Get test">
<br><br>

`routes/api.js`
```js
const players = [
  {firstName:"eli", lastName:"manning", position:"qb", age:37, team:"nyg"},
  {firstName:"tom", lastName:"brady", position:"qb", age:41, team:"nep"},
  {firstName:"jj", lastName:"watt", position:"de", age:28, team:"hou"}
]
const teams = [
  {name:"giants", city:"new york", conference:"nfc"},
  {name:"patriots", city:"new england", conference:"afc"},
  {name:"texans", city:"houston", conference:"afc"}
]


router.get('/player', (req, res) => {
  res.json({
    confirmation: 'success',
    data: players
  })
})

router.get('/team', (req, res) => {
  res.json({
    confirmation: 'success',
    data: teams
  })
})
```

<img src="./images/03.nodejs_api_get_player.png" width="70%" alt="REST API Get test">
<img src="./images/04.nodejs_api_get_team.png" width="70%" alt="REST API Get test">
<br><br>

위의 두 개의 분리된 엔드포인트를 request parameter인 resource 변수를 사용하여 더 일반적인 하나의 엔드포인트로 결합시킬 수 있다.

`routes/api.js`
```js
router.get('/:resource', (req, res) => {
  const resource = req.params.resource
  
  if (resource == 'team') {
    res.json({
      confirmation: 'success',
      data: teams
    })
  }

  if (resource == 'player') {
    res.json({
      confirmation: 'success',
      data: players
    })
  }
})
```

<img src="./images/05.nodejs_api_get_undefined.png" width="70%" alt="REST API Get test">
<br><br>

다음과 같이 더 간결하게 작성할 수 있다.

```js
const db = {
  teams: teams,
  player: players
}

router.get('/:resource', (req, res) => {
  const resource = req.params.resource

  const data = db[resource]
  if (data == null) {
    res.json({
      confirmation: 'fail',
      message: 'Invalid request. Resource undefined.'
    })
    return
  }
  res.json({
    confirmation: 'success',
    data: data
  })
})
```

## Mongo DB Connection
To be continued..

## Reference
* https://www.udemy.com/course/create-a-rest-api-with-node-js-and-mongo-db/