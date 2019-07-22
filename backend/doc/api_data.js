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
            "description": "<p>see <a href=\"#api-Service-Auth_Auth\">Auth service</a></p>"
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
        ],
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ]
      }
    },
    "filename": "api/auth/main.go",
    "groupTitle": "Auth"
  },
  {
    "type": "post",
    "url": "/avatar",
    "title": "AddAvatar",
    "version": "1.0.0",
    "group": "Avatar",
    "permission": [
      {
        "name": "self/admin"
      }
    ],
    "name": "AddAvatar",
    "description": "<p>Add avatar</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-Avatar_Create\">Avatar Service</a> <br> Max size is 5M</p>"
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
            "description": "<p>see <a href=\"#api-Service-Avatar_Create\">Avatar Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/avatar/main.go",
    "groupTitle": "Avatar",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
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
    "url": "/content",
    "title": "AddContent",
    "version": "1.0.0",
    "group": "Content",
    "permission": [
      {
        "name": "user/admin"
      }
    ],
    "name": "AddContent",
    "description": "<p>Add sell info content</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-Content_Create\">Content Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-Content_Create\">Content Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/content/main.go",
    "groupTitle": "Content",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "SellInfoServiceDown",
            "description": "<p>SellInfo service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "delete",
    "url": "/content",
    "title": "DeleteContent",
    "version": "1.0.0",
    "group": "Content",
    "permission": [
      {
        "name": "user/admin"
      }
    ],
    "name": "DeleteContent",
    "description": "<p>Delete sell info content</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-Content_Delete\">Content Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-Content_Delete\">Content Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/content/main.go",
    "groupTitle": "Content",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "SellInfoServiceDown",
            "description": "<p>SellInfo service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "get",
    "url": "/file/:fileId",
    "title": "GetFile",
    "version": "1.0.0",
    "group": "File",
    "permission": [
      {
        "name": "none"
      }
    ],
    "name": "GetFile",
    "description": "<p>Get file</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-File_Query\">File Service</a></p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "File",
            "optional": false,
            "field": "file",
            "description": "<p>file itself</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 404": [
          {
            "group": "Error 404",
            "optional": false,
            "field": "FileNotFound",
            "description": "<p>file not found</p>"
          }
        ],
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "FileServiceDown",
            "description": "<p>File service down</p>"
          }
        ]
      }
    },
    "filename": "api/file/main.go",
    "groupTitle": "File"
  },
  {
    "type": "post",
    "url": "/sellInfo",
    "title": "AddSellInfo",
    "version": "1.0.0",
    "group": "SellInfo",
    "permission": [
      {
        "name": "self/admin"
      }
    ],
    "name": "AddSellInfo",
    "description": "<p>Add sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-SellInfo_Create\">SellInfo Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-SellInfo_Create\">SellInfo Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/sellinfo/main.go",
    "groupTitle": "SellInfo",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "SellInfoServiceDown",
            "description": "<p>SellInfo service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "get",
    "url": "/sellInfo",
    "title": "FindSellInfo",
    "version": "1.0.0",
    "group": "SellInfo",
    "permission": [
      {
        "name": "none"
      }
    ],
    "name": "FindSellInfo",
    "description": "<p>Find sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-SellInfo_Find\">SellInfo Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-SellInfo_Find\">SellInfo Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/sellinfo/main.go",
    "groupTitle": "SellInfo",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "SellInfoServiceDown",
            "description": "<p>SellInfo service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "get",
    "url": "/sellInfo/:sellInfoId",
    "title": "GetSellInfo",
    "version": "1.0.0",
    "group": "SellInfo",
    "permission": [
      {
        "name": "none"
      }
    ],
    "name": "GetSellInfo",
    "description": "<p>Get sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-SellInfo_Query\">SellInfo Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-SellInfo_Query\">SellInfo Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/sellinfo/main.go",
    "groupTitle": "SellInfo",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "SellInfoServiceDown",
            "description": "<p>SellInfo service down</p>"
          }
        ]
      }
    }
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "Auth.Auth",
    "version": "1.0.0",
    "group": "Service",
    "name": "Auth_Auth",
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
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for invalid code <br> 3 for frozen user</p>"
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
    "filename": "srv/auth/main.go",
    "groupTitle": "Service"
  },
  {
    "type": "rpc",
    "url": "/rpc",
    "title": "Avatar.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "Avatar_Create",
    "description": "<p>Create avatar and return avatarId.</p>",
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
            "type": "bytes",
            "optional": false,
            "field": "file",
            "description": "<p>accept <a href=\"https://github.com/h2non/filetype#image\">file type</a></p>"
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
            "description": "<p>-1 for invalid param <br> 1 for success <br> 2 for not found <br> 3 for invalid file type</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "avatarId",
            "description": "<p>new avatar id</p>"
          }
        ]
      }
    },
    "filename": "srv/avatar/main.go",
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
    "title": "Content.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "Content_Create",
    "description": "<p>create sell info content</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "contentId",
            "description": "<p>24 bytes content id, left empty for first upload</p>"
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
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "contentId",
            "description": "<p>24 bytes contentId</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "contentToken",
            "description": "<p>random uuid content token</p>"
          }
        ]
      }
    },
    "filename": "srv/content/main.go",
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
    "title": "Content.Delete",
    "version": "1.0.0",
    "group": "Service",
    "name": "Content_Delete",
    "description": "<p>delete sell info content</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "contentId",
            "description": "<p>24 bytes content id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
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
          }
        ]
      }
    },
    "filename": "srv/content/main.go",
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
    "title": "File.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "File_Create",
    "description": "<p>Create file</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "bytes",
            "optional": false,
            "field": "file",
            "description": "<p>file stream bytes</p>"
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
            "description": "<p>-1 for invalid param <br> 1 for success</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "fileId",
            "description": "<p>file id</p>"
          }
        ]
      }
    },
    "filename": "srv/file/main.go",
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
    "title": "File.Delete",
    "version": "1.0.0",
    "group": "Service",
    "name": "File_Delete",
    "description": "<p>Delete file</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "fileId",
            "description": "<p>file id.</p>"
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
          }
        ]
      }
    },
    "filename": "srv/file/main.go",
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
    "title": "File.Query",
    "version": "1.0.0",
    "group": "Service",
    "name": "File_Query",
    "description": "<p>Query file stream</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "fileId",
            "description": "<p>file id.</p>"
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
            "type": "bytes",
            "optional": false,
            "field": "file",
            "description": "<p>file stream</p>"
          },
          {
            "group": "Success 200",
            "type": "int64",
            "optional": false,
            "field": "size",
            "description": "<p>file size</p>"
          }
        ]
      }
    },
    "filename": "srv/file/main.go",
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
    "title": "SellInfo.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "SellInfo_Create",
    "description": "<p>create sell info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>sellinfo userid</p>"
          },
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
            "optional": true,
            "field": "description",
            "description": "<p>description for good</p>"
          },
          {
            "group": "Parameter",
            "type": "double",
            "optional": true,
            "field": "price",
            "description": "<p>good price</p>"
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
    "filename": "srv/sellinfo/main.go",
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
    "title": "SellInfo.Find",
    "version": "1.0.0",
    "group": "Service",
    "name": "SellInfo_Find",
    "description": "<p>Find SellInfo.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": true,
            "field": "userId",
            "description": "<p>userId</p>"
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
            "field": "sellInfo",
            "description": "<p>see <a href=\"#api-Service-sellinfo_SellInfo_Query\">SellInfo Service</a></p>"
          }
        ]
      }
    },
    "filename": "srv/sellinfo/main.go",
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
    "title": "SellInfo.Query",
    "version": "1.0.0",
    "group": "Service",
    "name": "SellInfo_Query",
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
            "field": "sellInfoId",
            "description": "<p>sellInfoId</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "status",
            "description": "<p>1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired</p>"
          },
          {
            "group": "Success 200",
            "type": "int64",
            "optional": false,
            "field": "releaseTime",
            "description": "<p>sellInfo release time</p>"
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
            "type": "double",
            "optional": false,
            "field": "price",
            "description": "<p>good price</p>"
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
            "type": "int32",
            "optional": false,
            "field": "userId",
            "description": "<p>userId</p>"
          }
        ]
      }
    },
    "filename": "srv/sellinfo/main.go",
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
    "title": "User.Create",
    "version": "1.0.0",
    "group": "Service",
    "name": "User_Create",
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
            "type": "Response",
            "optional": false,
            "field": "user",
            "description": "<p>see <a href=\"#api-Service-user_User_Query\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "srv/user/main.go",
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
    "title": "User.Find",
    "version": "1.0.0",
    "group": "Service",
    "name": "User_Find",
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
    "filename": "srv/user/main.go",
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
    "title": "User.Query",
    "version": "1.0.0",
    "group": "Service",
    "name": "User_Query",
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
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "role",
            "description": "<p>user role, 1 for user <br> 2 for admin</p>"
          }
        ]
      }
    },
    "filename": "srv/user/main.go",
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
    "title": "User.Update",
    "version": "1.0.0",
    "group": "Service",
    "name": "User_Update",
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
            "optional": true,
            "field": "status",
            "description": "<p>user status, 1 for normal <br> 2 for frozen</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": true,
            "field": "role",
            "description": "<p>user role, 1 for user <br> 2 for admin</p>"
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
    "filename": "srv/user/main.go",
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
    "type": "post",
    "url": "/user",
    "title": "AddUser",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "name": "AddUser",
    "description": "<p>Add user, use default value.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "--",
            "optional": false,
            "field": "Param",
            "description": "<p>see <a href=\"#api-Service-User_Create\">User Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-User_Create\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
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
            "description": "<p>see <a href=\"#api-Service-User_Find\">User Service</a> <br> No param need admin permission!</p>"
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
            "description": "<p>see <a href=\"#api-Service-User_Find\">User Service</a> <br> None - studentId: hidden <br> None - studentName: hidden</p>"
          }
        ]
      }
    },
    "filename": "api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
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
        "name": "none/self/admin"
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
            "description": "<p>see <a href=\"#api-Service-User_Query\">User Service</a></p>"
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
            "description": "<p>see <a href=\"#api-Service-User_Query\">User Service</a> <br> studentId: hidden <br> studentName: hidden</p>"
          }
        ],
        "Self/Admin - Success 200": [
          {
            "group": "Self/Admin - Success 200",
            "type": "Response",
            "optional": false,
            "field": "response",
            "description": "<p>see <a href=\"#api-Service-User_Query\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
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
    "type": "put",
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
            "description": "<p>see <a href=\"#api-Service-User_Update\">User Service</a> <br> self not allow edit StudentId,StudentName,Status,Role</p>"
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
            "description": "<p>see <a href=\"#api-Service-User_Update\">User Service</a></p>"
          }
        ]
      }
    },
    "filename": "api/user/main.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
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
  }
] });
