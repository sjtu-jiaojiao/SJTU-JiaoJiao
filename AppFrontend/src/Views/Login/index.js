import React, { Component } from 'react';
import { WebView } from "react-native-webview";
import config from "../../Config";

export default class LoginScreen extends Component {
    render () {
        return <WebView source={{ uri: (config.fetchPrefix + 'v1/auth?code=bf53dbc0cdfb443fae1571480b4a0d87') }} />;
    }
}