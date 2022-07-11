### Backend APIs

Admin Login

    POST /admin/login HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 55
    
    {
        "name": "test",
        "password": "test_password"
    }

Admin Detail

    POST /admin/detail HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 2
    
    {}

Admin Register

    POST /admin/register HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 56
    
    {
        "name": "c_name",
        "password": "new password"
    }

Create Event

    POST /admin/create_event HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 396
    
    {
        "event": {
            "admin_id": 2,
            "name" : "an event",
            "description": "desp",
            "max_vote_num_per_person": 2,
            "candidates" : [
                {
                "name": "ca",
                "description": "niu"
                },
                {
                "name": "cb",
                "description": "bi"
                }
            ]
        },
        "voters": ["email1"]
    }

Get Event

    POST /admin/get_event HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 22
    
    {
        "event_id": 14
    }

End Event

    POST /admin/end_event HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cookie: adminId=1
    Content-Length: 22
    
    {
        "event_id": 13
    }
