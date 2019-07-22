import React, { Component } from 'react';
import {Alert, Text, TextInput, View, StyleSheet} from 'react-native';
import {Avatar, Button, ButtonGroup, SearchBar} from 'react-native-elements';
import Config from '../../Config';
import {NavigationActions} from "react-navigation";
import Textarea from 'react-native-textarea';

export default class ReleaseScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0,
            goodName: '',
            Discription: '',
            Price: '',
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

    updateDiscription = (Discription) => {
        this.setState({
            Discription: Discription,
        })
    };

    updatePrice = (Price) => {
        this.setState({
            Price: Price,
        })
    }

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
                            placeholder={'必填。输入物品名称'}
                            onChangeText={ this.updateGoodName }
                            value={ this.state.goodName }
                        />
                    </View>
                    <View style={{height: 10}} />
                    <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                        <Button
                            title='详细描述'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1.05}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            onPress={() => {alert(this.state.Discription)}}
                        />
                        <Textarea
                            containerStyle={{
                                flex: 4,
                                height: 100,
                                backgroundColor: 'white',
                                borderWidth: 0, //no effect
                                shadowColor: 'white', //no effect
                                borderBottomColor: 'transparent',
                                borderTopColor: 'transparent',
                            }}
                            style={{marginLeft: 14, marginRight: 13, fontSize: 18, }}
                            onChangeText={this.updateDiscription}
                            defaultValue={this.state.goodDiscription}
                            maxLength={360}
                            placeholder={'选填。可输入对物品的详细描述或者列出物品清单'}
                            placeholderTextColor={'grey'}
                            underlineColorAndroid={'transparent'}
                        />
                    </View>
                    <View style={{height: 10}} />
                    <View>
                        <Text style={{alignItems: 'center', justifyContent: 'center'}}>
                            商品标签
                        </Text>
                    </View>
                    <View style={{height: 10}} />
                    <View>
                        <Text style={{alignItems: 'center', justifyContent: 'center'}}>
                            上传照片和视频
                        </Text>
                    </View>
                    <View style={{height: 10}} />
                    <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                        <Button
                            title='预期价格'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1.1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            onPress={() => {alert(this.state.Price)}}
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
                            placeholder={'选填'}
                            onChangeText={ this.updatePrice }
                            value={ this.state.Price }
                        />
                    </View>
                    <View style={{height: 10}} />
                    <Button
                        title={'发布信息'}
                        titleStyle={{color: 'white', fontSize: 18}}
                        buttonStyle={{backgroundColor: 'red'}}
                        containerStyle={{width: 160, marginLeft: 120}}
                        raised={true}
                        onPress={() => {
                            // 检查必填项有没有写
                            // 检查price是不是最多两位的小数double，大于零
                            Alert.alert('你要发布的是**信息，确认发布嘛？');
                        }}
                    />
                    <View style={{height: 10}} />
                </View>
            )
        }
        //return (<View/>);
    }
}
