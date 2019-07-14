import React, { Component } from 'react';
import {Text, View, FlatList, Alert} from 'react-native';
import {SearchBar, Button, ListItem} from "react-native-elements";
import Config from '../../Config';
import {NavigationActions} from "react-navigation";

export default class SearchScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            searchText: '',
            searchType: 'User',
            showSearchItem: false,
            UserList: null,
        }
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}> </Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => (
        <ListItem
            bottomDivider
            title={item.title}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    );

    updateSearch = (searchText) => {
        this.setState({ searchText });
    };

    cancelSearch = (searchText) => {
        this.setState({
            showSearchItem: false,
            searchText: '',
        });
    };

    render() {
        const { searchText } = this.state.searchText;
        if (this.state.UserList !== null) {
            return (
                <View>
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            placeholder={'搜索'}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜索'
                            titleStyle={{color: 'white', fontSize: 17}}
                            buttonStyle={{backgroundColor: 'red'}}
                            containerStyle={{height: 50}}
                            raised={true}
                            onPress={() => {
                                fetch((Config.fetchPrefix + 'user/?userName=' + this.state.searchText))
                                    .then((response) => {response.json();console.warn(responseJson);})
                                    .then((responseJson) => {

                                        this.setState({
                                            UserList: responseJson.user,
                                        });
                                        //console.warn(this.state.UserList);
                                        //console.warn(responseJson);
                                    })
                                    .catch((error) => console.error(error));
                            }}
                        />
                    </View>
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.UserList}
                        renderItem={this.renderItem}
                    />
                </View>
            )
        }
        else {
            return (
                <View>
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            containerStyle={{width: 340}}
                            placeholder={'搜索'}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜索'
                            titleStyle={{color: 'white', fontSize: 17}}
                            buttonStyle={{backgroundColor: 'steelblue'}}
                            containerStyle={{height: 50}}
                            raised={true}
                            onPress={() => {
                                if(this.state.searchText !== '') {
                                    console.warn(this.state.searchText);
                                    //fetch((Config.fetchPrefix + 'user/?userName=' + this.state.searchText))
                                    fetch('http://202.120.40.8:30711/v1/user?userName=c')
                                        .then((response) => {
                                            console.warn(response);
                                            response.json();
                                        })
                                        .then((responseJson) => {
                                            this.setState({
                                                UserList: responseJson.user,
                                            });
                                            console.warn(this.state.UserList);
                                            //console.warn(responseJson);
                                        })
                                        .catch((error) => console.error(error));
                                }
                            }}
                        />
                    </View>
                </View>
            )
        }
    }
}
