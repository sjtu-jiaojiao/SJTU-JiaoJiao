import React, { Component } from 'react';
import { Text, View } from 'react-native';
import {SearchBar, CheckBox  } from "react-native-elements";

export default class SearchScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            searchText: '',
            searchType: 'Sell',
        }
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}></Text>)
    };

    updateSearch = (searchText) => {
        this.setState({ searchText });
    };

    render() {
        const { searchText } = this.state.searchText;
        return (
            <View>
                <SearchBar
                    placeholder={'搜索'}
                    onChangeText={ this.updateSearch }
                    value={ this.state.searchText }
                />
            </View>
        )
    }
}