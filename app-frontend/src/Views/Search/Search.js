import React, { Component } from 'react';
import {Text, View, FlatList, Alert, TextInput} from 'react-native';
import {SearchBar, Button, ListItem, ButtonGroup, Avatar} from "react-native-elements";
import Config from '../../Config';

let dev = "http://202.120.40.8:30711/v1";

class Http {
    // 静态方法
    static get(url, params) {
        // 将后台接口的公共部分拼接进去
        url = dev + url;
        //判断有木有参数
        if (params) {
            // 定一个空数组
            let paramsArray = [];
            //  拆分对象
            Object.keys(params).forEach(key =>
                paramsArray.push(key + "=" + params[key])
            );
            // 判断是否地址拼接的有没有 ？,当没有的时候，使用 ？拼接第一个参数，如果有参数拼接，则用&符号拼接后边的参数
            if (url.search(/\?/) === -1) {
                url = url + "?" + paramsArray.join("&");
            } else {
                url = url + "&" + paramsArray.join("&");
            }
        }
        // 返回一个promise
        return new Promise((resolve, reject) => {
            //fetch请求
            fetch(url, { method: "GET" })
                .then(response => response.json())
                .then(resulet => {
                    resolve(resulet);
                })
                .catch(error => {
                    reject(error);
                });
        });
    }
}

export default class SearchScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            searchText: '',
            DataList: null,
            selectedIndex: 1,
        };
        this.updateIndex = this.updateIndex.bind(this);
    };

    updateIndex (selectedIndex) {
        this.setState({selectedIndex})
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}> </Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => (
        <ListItem
            bottomDivider
            title={() => {
                if(this.state.selectedIndex === 0)
                    return item.userName;
                else if (this.state.selectedIndex === 1)
                    return item.goodName;
                else if (this.state.selectedIndex === 2)
                    return item.goodName;
            }}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    );

    updateSearch = (searchText) => {
        this.setState({ searchText: searchText });
    };

    cancelSearch = (searchText) => {
        this.setState({
            searchText: '',
            DataList: null,
        });
    };

    startSearch = () => {
        let {DemoUser} = [
            {
                userName: this.state.searchText,
            },
            {
                userName: this.state.selectedIndex,
            },
            {
                userName: 'JBoss',
            },
        ];

        let {DemoGoods} = [
            {
                goodName: this.state.searchText,
            },
            {
                goodName: this.state.selectedIndex,
            },
            {
                goodName: 'LBoss',
            },
        ];

        console.warn('Search ' + this.state.searchText + ' in ' + this.state.selectedIndex);
        if (this.state.selectedIndex === 0) {
            //console.warn('Search User');
            this.setState({
                DataList: DemoUser,
            }, () => {
                /* Empty */
            });
        } else {
            //console.warn('Search goods');
            this.setState({
                DataList: DemoGoods,
            }, () => {
                /* Empty */
            });
        }
        /*if(this.state.searchText !== '') {
            let obj = { userName: this.state.searchText };
            Http.get('/user', obj)
                .then((response) => {
                    this.setState({
                        UserList: response.user,
                    });
                    //console.warn(this.state.UserList);
                })
                .catch((error) => console.error(error));
        }*/
        // 注意，如果返回的list为空，要提示用户没有他想要的！
    };

    render() {
        const buttons = ['用户', '出售信息', '求购信息'];
        const { selectedIndex } = this.state;
        const { searchText } = this.state;
        if (this.state.DataList !== null) {
            console.warn('Have DataList Now!');
            console.warn(this.state.DataList);
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
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            searchIcon={false}
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
                            placeholder={'搜索' + buttons[this.state.selectedIndex]}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            clearIcon={
                                <Avatar
                                    rounded size='small'
                                    source={require('../../assets/icons/cancel.png')}
                                />
                            }
                            cancelIcon={false}
                            onPress={() => this.startSearch()}
                        />
                    </View>
                    <View style={{height: 10}} />
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.DataList}
                        renderItem={this.renderItem}
                    />
                </View>
            )
        }
        else {
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
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            searchIcon={false}
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
                            placeholder={'搜索' + buttons[this.state.selectedIndex]}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            clearIcon={
                                <Avatar
                                    rounded size='small'
                                    source={require('../../assets/icons/cancel.png')}
                                />
                            }
                            cancelIcon={false}
                            onPress={() => this.startSearch()}
                        />
                    </View>
                    <View style={{height: 10}} />
                </View>
            )
        }
    }
}
