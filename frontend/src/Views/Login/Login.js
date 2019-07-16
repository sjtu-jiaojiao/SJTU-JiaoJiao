import React, { Component } from 'react';
import { Text, View } from 'react-native';
import {WebView} from "react-native-webview";
import {NavigationActions} from "react-navigation";
import Config from "../../Config";

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
                    //console.warn(Config.JaccountToken.status);
                    fetch((Config.fetchPrefix + 'user/1'),{
                        headers: {
                            Authorization: ('Bearer ' + Config.JaccountToken.token),
                        }
                    })
                        .then((response) => {
                            //console.warn(response);
                            return response.json();
                        })
                        .then((responseJson) => {
                            Config.userInfo.userId=responseJson.userId;
                            Config.userInfo.userName=responseJson.userName;
                            Config.userInfo.studentId=responseJson.studentId;
                            Config.userInfo.studentName=responseJson.studentName;
                            Config.userInfo.telephone=responseJson.telephone;
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
