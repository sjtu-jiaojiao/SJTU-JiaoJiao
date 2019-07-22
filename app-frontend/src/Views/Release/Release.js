import React, { Component } from 'react';
import {Alert, Text, TextInput, View} from 'react-native';
import {Avatar, Button, ButtonGroup, SearchBar} from 'react-native-elements';
import Config from '../../Config';
import {NavigationActions} from "react-navigation";

export default class ReleaseScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0,
            goodName: '',
        };
        this.updateIndex = this.updateIndex.bind(this);
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>发布交易信息</Text>)
    };

    updateIndex (selectedIndex) {
        this.setState({
            selectedIndex: selectedIndex,
        })
    }

    updateGoodName = (goodName) => {
        this.setState({
            goodName: goodName,
        })
    };

    render() {
        const buttons = ['出售信息', '求购信息'];
        const { selectedIndex } = this.state;
        /*
        if (Config.userInfo.userId === -1) {
            Alert.alert(
                '未登录',
                '无法在未登录状态下发布信息，是否切换至登录界面？',
                [
                    {
                        text: '取消',
                        onPress: () => {
                            this.props.navigation.navigate('Home');
                        },
                        style: 'cancel',
                    },
                    {
                        text: '确定', onPress: () => {
                            this.props.navigation.navigate('Login');
                        }
                    },
                ],
                {cancelable: false},
            )
        } else */{
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <ButtonGroup
                        onPress={this.updateIndex}
                        selectedIndex={selectedIndex}
                        buttons={buttons}
                        containerStyle={{height: 35}}
                        disabledTextStyle={{fontSize: 17}}
                        selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                    />
                    <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                        <Button
                            title='物品名称'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1.1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            onPress={() => {alert(this.state.goodName)}}
                        />
                        <SearchBar
                            searchIcon={false}
                            clearIcon={false}
                            cancelIcon={false}
                            lightTheme={true}
                            containerStyle={{
                                flex: 4,
                                height: 55,
                                backgroundColor: 'white',
                                borderWidth: 0, //no effect
                                shadowColor: 'white', //no effect
                                borderBottomColor: 'transparent',
                                borderTopColor: 'transparent',
                            }}
                            inputContainerStyle={{
                                marginTop: -9,
                                height: 55,
                                backgroundColor: 'white',
                            }}
                            placeholder={'输入' + buttons[this.state.selectedIndex] + '物品名称'}
                            onChangeText={ this.updateGoodName }
                            value={ this.state.goodName }
                        />
                    </View>
                    <View style={{height: 10}} />
                </View>
            )
        }
        //return (<View/>);
    }
}
