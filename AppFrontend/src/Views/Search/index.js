import React, { Component } from 'react';
import { Text, View } from 'react-native';
import {SearchBar} from "react-native-elements";

export default class SearchScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            search: '',
        }
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}></Text>)
    };

    updateSearch = (search) => {
        this.setState({ search });
    };

    render() {
        const { search } = this.state.search;
        return (
            <View>
                <View>
                    <SearchBar
                        placeholder={'搜索'}
                        onChangeText={ this.updateSearch }
                        value={ this.state.search }
                    />
                </View>
            </View>
        )
    }
}