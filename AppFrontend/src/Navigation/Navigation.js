import React, { Component } from 'React';
import { View, Text} from 'react-native';
import {
    createBottomTabNavigator,
    createStackNavigator,
    createAppContainer
} from "react-navigation";

import HomeScreen from '../Views/Home'
import ReleaseScreen from '../Views/Release'
import ContactScreen from '../Views/Contact'
import UserScreen from '../Views/User'

class Test extends Component {
    render() {
        return (
            <View >
                <Text>
                    This is a test page!
                </Text>
            </View>
        )
    }
}

const HomeStack = createStackNavigator({
    Home: HomeScreen,
    Test: Test,
});

const ReleaseStack = createStackNavigator({
    Release: ReleaseScreen,
    Test: Test,
});

const ContactStack = createStackNavigator({
    Contact: ContactScreen,
    Test: Test,
});

const UserStack = createStackNavigator({
    User: UserScreen,
    Test: Test,
});

const TabBar = createBottomTabNavigator({
    Home: HomeStack,
    Release: ReleaseStack,
    Contact: ContactStack,
    User: UserStack,
},{
    /* 目前没有其他配置先 */
});

const AppContainer = createAppContainer(TabBar);

export default AppContainer;