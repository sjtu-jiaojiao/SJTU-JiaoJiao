define({ "api": [
  {
    "type": "get",
    "url": "/auth",
    "title": "GetUser",
    "version": "0.1.0",
    "group": "User",
    "name": "Auth",
    "description": "<p>This is the Description.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>Firstname of the User.</p>"
          }
        ]
      }
    },
    "filename": "./api/auth/main.go",
    "groupTitle": "User"
  }
] });
