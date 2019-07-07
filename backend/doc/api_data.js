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
    "description": "<p>Redirect to OAuth url.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "code",
            "description": "<p>OAuth code callback, DO NOT call it by yourself.</p>"
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
            "description": "<p>Redirect to OAuth url.</p>"
          }
        ],
        "With param - Success 200": [
          {
            "group": "With param - Success 200",
            "type": "--",
            "optional": false,
            "field": "Response",
            "description": "<p>see <a href=\"#api-Service-auth_Auth_Auth\">Auth service</a></p>"
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
            "type": "String",
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
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>1 for success, 2 for empty code, 3 for invalid code</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
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
