# Event log API

API uses unix epoch ('from' and 'to' params) as timestamps.

**API endpoints**:
1. GET /api/v1/events/type:\<type\>
2. GET /api/v1/events/from:\<from\>/to:\<to\>
3. GET /api/v1/events/type:\<type\>/from:\<from\>/to:\<to\>
4. POST /api/v1/event

**Start environment**: docker-compose up -d

**Examples** _(curl or postman)_:

1. Get events by type example: curl localhost:3000/api/v1/events/type:user_click 
2. Get events by time range example: curl localhost:3000/api/v1/events/from:1622575864/to:1622575866
3. Get events by type and time range example: curl localhost:3000/api/v1/events/type:user_click/from:1622575864/to:1622575866
4. Add event example: curl --header "Content-Type: application/json" --request POST --data '{"event_timestamp":"1625053791","event_type":"user_logout","event_content":"Test content for user logout"}' localhost:3000/api/v1/event


_TODO: add mocks and tests_