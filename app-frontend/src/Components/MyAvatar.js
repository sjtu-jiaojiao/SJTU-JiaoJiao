import {Avatar} from "react-native-elements";
import React, {Component} from 'react';
import Config from "../Config";

export default class MyAvatar extends Component {
    render() {
        if(this.props.avatarID === 'to_be_changed') {
            return (
                <Avatar rounded size='large' source={require('../assets/images/NotLogin.jpg')}/>
            )
        } else {
            return (
                <Avatar rounded size='large' source={{uri: (Config.fetchPrefix + 'file/' + this.props.avatarID)}}/>
            )
        }
    }
}
