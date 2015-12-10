<p align="center" style="text-align: center">
  <img align="center" src="https://s3.amazonaws.com/cdn.catarse/assets/catartico72.png"/>
</p>
<h3 align="center">Catartico</h3>
<br/>

**Catartico** is a simple [slack incomming webhooks](https://api.slack.com/incoming-webhooks) wrapper writens in Go thats trigger 
the post message using PostgreSQL ```LISTEN / NOTIFY```.


#### How I can use?


```
go get gitub.com/devton/catartico

SLACK_HOOK_URL=https://hooks.slack.com/services/TOKEN \
  DB_HOST=localhost \
  DB_USER=username \
  DB_PORT=5432
  DB_PASSWORD=dbpass \
  DB_DATABASE=catarse_development \
  LISTEN_BUCKET=catartico_events \
  catartico
```


You can put these ENV vars into a file an run only ```catartico```


#### Testing with NOTIFY

Asuming that you have the catarse project database populated whe can peform a NOTIFY like:

```sql
SELECT 
    pg_notify(
        'catartico_events', json_build_object(
            'channel', '#some_channel_to_post_on_slack', -- can use @username or #channel
            'name', 'event_name', -- event name
            'title', 'hello this is a *title* can use markdown', -- title of message
            'text', 'this is a text')::text -- text of message
    );
```

You will see on console something like this:
![](https://s3.amazonaws.com/catarse.files/Captura+de+Tela+2015-12-10+às+10.24.22+AM.png)

And on slack the message:

![](https://s3.amazonaws.com/catarse.files/Captura+de+Tela+2015-12-10+às+6.39.14+PM.png)

o/





