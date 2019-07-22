import React, { Component } from "react";
import {Alert, Text, View, TextInput} from "react-native";
import { Button, Overlay, ButtonGroup } from "react-native-elements";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";


export default class ContactScreen extends Component {
    constructor () {
        super()
        this.state = {
            selectedIndex: 1
        };
        this.updateIndex = this.updateIndex.bind(this);
    }

    updateIndex (selectedIndex) {
        this.setState({selectedIndex})
    }

    render () {
        const buttons = ['用户', '出售信息', '求购信息'];
        const { selectedIndex } = this.state;

        return (
            <View>
                <ButtonGroup
                    onPress={this.updateIndex}
                    selectedIndex={selectedIndex}
                    buttons={buttons}
                    containerStyle={{height: 30}}
                />
                <Text>
                    {this.state.selectedIndex}
                </Text>
                <Text>
                    {buttons[this.state.selectedIndex]}
                </Text>
            </View>
        )
    }
}

