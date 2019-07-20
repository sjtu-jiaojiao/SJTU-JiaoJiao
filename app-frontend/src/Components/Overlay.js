import React, { Component } from "react";
import {Alert, Text, View, TextInput} from "react-native";
import { Button, Overlay } from "react-native-elements";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";


export default class Overlay extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isVisible: false,
        };
        this.inputValue = '';
    }

    render() {
        let updateType = this.props.updateType
        return (
            <View>
                <Overlay
                    isVisible={this.state.isVisible}
                    windowBackgroundColor="rgba(255, 255, 255, .5)"
                    overlayBackgroundColor='#EFEFF5'
                    width="auto"
                    height="auto"
                    onBackdropPress={() => this.setState({ isVisible: false })}
                >
                    <View style={{width: 300}}>
                        <TextInput
                            style={{borderWidth: 1, borderColor: '#cccccc', textAlign: 'center'}}
                            onChangeText={text => this.inputValue = text}
                        />
                        <View style={{flexDirection: 'row', justifyContent: 'space-around'}}>
                            <Button
                                title='取消'
                                titleStyle={{color: 'red', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => this.setState({isVisible: false})}
                            />
                            <Button
                                title='确认修改'
                                titleStyle={{color: '#298BFF', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => this.setState({ isVisible: false })}
                            />
                        </View>
                    </View>
                </Overlay>
                <Text>
                    这是信息界面
                </Text>
                <Button
                    title='show'
                    titleStyle={{color: 'white', fontSize: 17}}
                    buttonStyle={{backgroundColor: 'red'}}
                    containerStyle={{width: 160, marginLeft: 120}}
                    raised={true}
                    onPress={() => this.setState({isVisible: true})}
                />
            </View>
        );
    }
}

