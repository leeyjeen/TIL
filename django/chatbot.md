# Chatbot in Python, Django

## Install packages
**chatterbot**을 설치한다.
```bash
$ pip install chatterbot
```

### Installed Apps
django project ```settings.py``` 파일의 ```INSTALLED_APPS```에 ```chatterbot.ext.django_chatterbot```을 추가한다.
```python
INSTALLED_APPS = [
    # ...
    'chatterbot.ext.django_chatterbot',
]
```

### Migrations
chatbot을 위한 db migration을 한다.
```bash
$ python manage.py migrate django_chatterbot
```

### MongoDB and Django
ChatterBot은 MongoDB용 스토리지 어댑터를 가지고 있으나, Django에서 작동하지 않는다. Django와 챗봇을 위한 데이터베이스로 MongoDB를 사용하고 싶다면, Django MongoDB Engine과 같은 Django 스토리지 백엔드를 설치해야 할 것이다. 

이것이 필요한 이유는 Django의 스토리지 백엔드가 ChatterBot의 스토리지 어댑터와 다르고 완전히 분리되어 있기 때문이다.


## Chatterbot Django Settings
Django ```settings.py``` 파일에서 ChatterBot configuration을 편집한다.
```python
CHATTERBOT = {
    'name': 'Tech Supprot Bot',
    'logic_adaptera': [
        'chatterbot.logic.MathematicalEvaluation',
        'chatterbot.logic.TimeLogicAdapter',
        'chatterbot.logic.BestMatch'
    ]
}
```

### Additional Django settings
- ```django_app_name``` [default: 'django_chatterbot'] - 모델을 찾을 Djago 앱 이름 (?)


## ChatterBot Django Views
### Example API Views
다음 예시는 ChatterBot을 이용하여 애플리케이션의 대화형 API 엔드포인트를 생성하는 하나의 방법을 보여주는 API view를 제공한다.

엔드포인트는 다음과 같은 형식의 JSON request를 예상한다.
```json
{ "text": "My input statement" }
``` 

```python
class ChatterBotApiView(View):
    """
    Provide an API endpoint to interact with ChatterBot.
    """
    
    chatterbot = ChatBot(**settings.CHATTERBOT)

    def post(self, request, *args, **kwargs):
        """
        Return a response to the statement in the posted data.
        * The JSON data should contain a 'text' attribute.
        """
        input_data = json.loads(request.body.decode('utf-8'))
        
        if 'text' not in input_data:
            return JsonResponse({
                'text': [
                    'The attribute "text" is required.'
                ]
            }, status=400)
        response = self.chatterbot.get_response(input_data)

        response_data = response.serialize()

        return JsonResponse(response_data, status=200)

    def get(self, request, *args, **kwargs):
        """
        Return data corresponding to the current conversation.
        """
        return JsonResponse({
            'name': self.chatterbot.name
        })
```
전체 예시는 
https://github.com/gunthercox/ChatterBot/tree/master/examples/django_app 에서 확인할 수 있다.

## Webservices
Django app을 호스팅하려면 이를 호스팅할 방법을 선택해야 한다. Heroku, PythonAnyWhere 등의 무료 서비스를 이용할 수 있다.

### WSGI
Python 웹 애플리케이션을 제공하는 일반적인 방법은 WSGI(Web Server Gateway Interface) 패키지를 사용하는 것이다.

Gunicorn은 WSGI 서버를 위한 훌륭한 선택이다. 그들은 웹사이트에 상세한 문서와 설치 지침을 가지고 있다.

### Hosting static files
Django 애플리케이션에 정적 파일을 호스팅하는 많은 방법들이 있다. 이를 위한 특별히 쉬운 방법 중 하나는 웹 애플리케이션에서 정적 파일을 제공할 수 있도록 고안된 python 패키지인 **WhiteNoise**를 사용하는 것이다.

## Reference
- https://chatterbot.readthedocs.io/en/stable/
- https://chatterbot.readthedocs.io/en/stable/django/index.html