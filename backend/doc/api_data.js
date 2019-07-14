define({ "api": [
  {
    "type": "get",
    "url": "/auth",
    "title": "GetAuth",
    "version": "1.0.0",
    "group": "Auth",
    "permission": [
      {
        "name": "none"
      }
    ],
    "name": "GetAuth",
    "description": "<p>Redirect to OAuth url</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "code",
            "description": "<p>OAuth code callback, DO NOT call it by yourself</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "No param - Redirect 301": [
          {
            "group": "No param - Redirect 301",
            "type": "Redirect",
            "optional": false,
            "field": "url",
            "description": "<p>Redirect to OAuth url</p>"
          }
        ],
        "With param - Success 200": [
          {
            "group": "With param - Success 200",
            "type": "--",
            "optional": false,
            "field": "Response",
            "description": "<p>see <a href=\"#api-Service-auth_Auth_Auth\">Auth service</a> <br> status: -1 is not allowed</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "AuthServiceDown",
            "description": "<p>Auth service down</p>"
          }
        ]
      }
    },
    "filename": "./api/auth/main.go",
    "groupTitle": "Auth"
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "auth.Auth.Auth",
    "version": "1.0.0",
    "group": "Service",
    "name": "auth_Auth_Auth",
    "description": "<p>Check OAuth code.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "code",
            "description": "<p>OAuth code.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for invalid code</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "token",
            "description": "<p>verified token when status=1</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "OAuthServerDown",
            "description": "<p>can't connect to OAuth server</p>"
          }
        ]
      }
    },
    "filename": "./srv/auth/main.go",
    "groupTitle": "Service"
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "sellinfo.Content.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "sellinfo_Content_Create",
    "description": "<p>create sell info content</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "contentId",
            "description": "<p>content id, left empty for first upload</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "contentToken",
            "description": "<p>content token, left empty for first upload</p>"
          },
          {
            "group": "Parameter",
            "type": "bytes",
            "optional": false,
            "field": "content",
            "description": "<p>binary content</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "type",
            "description": "<p>1 for picture <br> 2 for video</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for invalid token</p>"
          }
        ]
      }
    },
    "filename": "./srv/sellinfo/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "sellinfo.SellInfo.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "sellinfo_SellInfo_Create",
    "description": "<p>create sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int64",
            "optional": false,
            "field": "validTime",
            "description": "<p>valid timestamp</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "goodName",
            "description": "<p>good name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>description for good</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "contentId",
            "description": "<p>content id of good</p>"
          },
          {
            "group": "Parameter",
            "type": "array",
            "optional": true,
            "field": "tag",
            "description": "<p>tags for good(un-finished)</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "contentToken",
            "description": "<p>content token</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for invalid token</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "sellInfoId",
            "description": "<p>created sellInfoId</p>"
          }
        ]
      }
    },
    "filename": "./srv/sellinfo/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "sellinfo.SellInfo.Query",
    "version": "1.0.0",
    "group": "Service",
    "name": "sellinfo_SellInfo_Query",
    "description": "<p>Query sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "sellInfoId",
            "description": "<p>sellInfo id.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for non-exist</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "sellInfoId",
            "description": "<p>sellInfoId</p>"
          },
          {
            "group": "Success 200",
            "type": "int64",
            "optional": false,
            "field": "validTime",
            "description": "<p>sellInfo validate time</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "goodName",
            "description": "<p>good name</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>good description</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "contentId",
            "description": "<p>multimedia data</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "tag",
            "description": "<p>good tags (un-finished)</p>"
          }
        ]
      }
    },
    "filename": "./srv/sellinfo/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.AdminUser.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_AdminUser_Create",
    "description": "<p>Create new admin user.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "studentId",
            "description": "<p>student id.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for exist user</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "adminId",
            "description": "<p>created or existed adminId</p>"
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.AdminUser.Find",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_AdminUser_Find",
    "description": "<p>Find admin user.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "studentId",
            "description": "<p>student id.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for not found</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "adminId",
            "description": ""
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.User.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_User_Create",
    "description": "<p>Create new user.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "studentId",
            "description": "<p>student id.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "studentName",
            "description": "<p>student name.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for exist user</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>created or existed userid</p>"
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.User.Find",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_User_Find",
    "description": "<p>Find user(fuzzy).</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "userName",
            "description": "<p>username</p>"
          },
          {
            "group": "Parameter",
            "type": "uint32",
            "optional": false,
            "field": "limit",
            "defaultValue": "100",
            "description": "<p>row limit</p>"
          },
          {
            "group": "Parameter",
            "type": "uint32",
            "optional": false,
            "field": "offset",
            "defaultValue": "0",
            "description": "<p>row offset</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "list",
            "optional": false,
            "field": "user",
            "description": "<p>see <a href=\"#api-Service-user_User_Query\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.User.Query",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_User_Query",
    "description": "<p>Query user info.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>user id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "userName",
            "description": "<p>user name</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "avatarId",
            "description": "<p>user avatar id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "telephone",
            "description": "<p>user telephone</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "studentId",
            "description": "<p>student id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "studentName",
            "description": "<p>student name</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>user status, 1 for normal <br> 2 for frozen</p>"
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "user.User.Update",
    "version": "1.0.0",
    "group": "Service",
    "name": "user_User_Update",
    "description": "<p>Update user info, only update provided field. If clearEmpty=1 and param support allow clear, clear the field when not provided.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>user id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "userName",
            "description": "<p>user name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "avatarId",
            "description": "<p>user avatar id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "telephone",
            "description": "<p>user telephone, allow clear</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "studentId",
            "description": "<p>student id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "studentName",
            "description": "<p>student name</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>user status, 1 for normal <br> 2 for frozen</p>"
          },
          {
            "group": "Parameter",
            "type": "bool",
            "optional": false,
            "field": "clearEmpty",
            "defaultValue": "0",
            "description": "<p>clear the empty field</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for user not found</p>"
          }
        ]
      }
    },
    "filename": "./srv/user/main.go",
    "groupTitle": "Service",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "DBServerDown",
            "description": "<p>can't connect to database server</p>"
          }
        ]
      }
    }
  },
  {
    "type": "get",
    "url": "/user",
    "title": "FindUser",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "none/admin"
      }
    ],
    "name": "FindUser",
    "description": "<p>Find user</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-user_User_Find\">User Service</a> <br> No param need admin permission!</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Response",
            "optional": false,
            "field": "response",
            "description": "<p>see <a href=\"#api-Service-user_User_Find\">User Service</a> <br> None - studentId: hidden <br> None - studentName: hidden</p>"
          }
        ]
      }
    },
    "filename": "./api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "UserServiceDown",
            "description": "<p>User service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "get",
    "url": "/user/:userId",
    "title": "GetUserInfo",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "none/self"
      }
    ],
    "name": "GetUserInfo",
    "description": "<p>Get user info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-user_User_Query\">User Service</a></p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "None - Success 200": [
          {
            "group": "None - Success 200",
            "type": "Response",
            "optional": false,
            "field": "response",
            "description": "<p>see <a href=\"#api-Service-user_User_Query\">User Service</a> <br> studentId: hidden <br> studentName: hidden</p>"
          }
        ],
        "Self - Success 200": [
          {
            "group": "Self - Success 200",
            "type": "Response",
            "optional": false,
            "field": "response",
            "description": "<p>see <a href=\"#api-Service-user_User_Query\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "./api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "UserServiceDown",
            "description": "<p>User service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "post",
    "url": "/user",
    "title": "UpdateUser",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "self/admin"
      }
    ],
    "name": "UpdateUser",
    "description": "<p>Update user</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-user_User_Update\">User Service</a></p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Response",
            "optional": false,
            "field": "response",
            "description": "<p>see <a href=\"#api-Service-user_User_Update\">User Service</a> <br> status: -1 is not allowed</p>"
          }
        ]
      }
    },
    "filename": "./api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "UserServiceDown",
            "description": "<p>User service down</p>"
          }
        ]
      }
    }
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_home_imwxz_Documents_code_SJTU_JiaoJiao_backend_doc_main_js",
    "groupTitle": "_home_imwxz_Documents_code_SJTU_JiaoJiao_backend_doc_main_js",
    "name": ""
  }
] });