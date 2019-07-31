import React, { Component } from 'react';
import { Text, View } from 'react-native';
import {WebView} from "react-native-webview";
import {NavigationActions} from "react-navigation";
import Config from "../../Config";
import jwt_decode from 'jwt-decode';

export default class LoginScreen extends Component {

    render() {
        const run = `
            let x=document.getElementsByTagName("pre")[0].childNodes[0].nodeValue;
            window.ReactNativeWebView.postMessage(x);
        `;
        return (
            <WebView
                ref={r => (this.webref = r)}
                source={{ uri: (Config.fetchPrefix + 'auth') }}
                onMessage={(event) => {
                    //console.warn(event.nativeEvent.data);
                    Config.JaccountToken = JSON.parse(event.nativeEvent.data);
                    //console.warn(Config.JaccountToken);
                    let decodeToken = jwt_decode(Config.JaccountToken.token);
                    //console.warn(decodeToken);
                    //console.warn(Config.JaccountToken);
                    //console.warn(Config.JaccountToken.status);
                    fetch((Config.fetchPrefix + 'user/' + decodeToken.id),{
                        headers: {
                            Authorization: ('Bearer ' + Config.JaccountToken.token),
                        }
                    })
                        .then((response) => {
                            //console.warn(response);
                            return response.json();
                        })
                        .then((responseJson) => {
                            Config.userInfo.userID=responseJson.userID;
                            Config.userInfo.userName=responseJson.userName;
                            Config.userInfo.studentID=responseJson.studentID;
                            Config.userInfo.studentName=responseJson.studentName;
                            Config.userInfo.telephone=responseJson.telephone;
                            Config.userInfo.avatarID=responseJson.avatarID;
                            //console.warn(Config.userInfo);
                            this.props.navigation.reset([NavigationActions.navigate({ routeName: 'User' })], 0);
                            //console.warn(responseJson);
                        })
                        .catch((error) => console.error(error));
                }}
                onNavigationStateChange={(event)=>{
                    if(event.title.includes('?code='))
                    {
                        // console.warn(event);
                        this.webref.injectJavaScript(run);
                    }
                }}
            />
        )
    }
}
