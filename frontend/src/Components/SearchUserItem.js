import React, { Component } from "react";
import { View, Text } from 'react-native';
import { ListItem } from "react-native-elements";

export default class SearchUserItem extends Component {
    render() {
        return (
            <View>
                {
                    list.map((item, i) => (
                        <ListItem
                            bottomDivider
                            key={i}
                            title={item.userName}
                            titleStyle={{ color: 'black', fontSize: 17 }}
                        />
                    ))
                }
            </View>
        );
    }
}