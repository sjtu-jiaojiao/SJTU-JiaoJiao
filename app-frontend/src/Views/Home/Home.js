import React, {Component} from 'react';
import { Text, View } from 'react-native';
import {Avatar, Button, Divider, ListItem } from "react-native-elements";
import Icon from "react-native-vector-icons/AntDesign";

export default class HomeScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>交大交交</Text>),
    };

    render() {
        return (
            <View style={{backgroundColor: '#EFEFF5'}}>
                <View style={{height: 20}} />
                <Button
                    icon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/search.png')}
                        overlayContainerStyle={{backgroundColor: '#EFEFF5'}}
                    />}
                    title=' 点我搜索'
                    titleStyle={{color: 'black', fontSize: 17}}
                    buttonStyle={{width: 120, backgroundColor: '#EFEFF5'}}
                    containerStyle={{width: 120, marginLeft: 20}}
                    raised={true}
                    onPress={() => this.props.navigation.navigate('Search')}
                />
                <View style={{height: 20}} />
                <ListItem
                    containerStyle={{height: 100}}
                    leftAvatar={<Avatar
                        rounded size='medium'
                        source={require('../../assets/icons/sellIcon.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    title={'出售信息'}
                    titleStyle={{ color: 'black', fontWeight: 'bold', fontSize: 20 }}
                    subtitle={'点击查看所有已经发布的出售信息'}
                    subtitleStyle={{ color: 'grey', fontSize: 14 }}
                    rightIcon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/right.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    onPress={() => this.props.navigation.navigate('SellInfo')}
                />
                <View style={{height: 20}} />
                <ListItem
                    containerStyle={{height: 100}}
                    leftAvatar={<Avatar
                        rounded size='medium'
                        source={require('../../assets/icons/buyIcon.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    title={'求购信息'}
                    titleStyle={{ color: 'black', fontWeight: 'bold', fontSize: 20 }}
                    subtitle={'点击查看所有已经发布的求购信息'}
                    subtitleStyle={{ color: 'grey', fontSize: 14 }}
                    rightIcon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/right.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    onPress={() => this.props.navigation.navigate('BuyInfo')}
                />
                <View style={{height: 20}} />
                <Button
                    icon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/search.png')}
                        overlayContainerStyle={{backgroundColor: '#EFEFF5'}}
                    />}
                    title=' TestPage'
                    titleStyle={{color: 'black', fontSize: 17}}
                    buttonStyle={{width: 120, backgroundColor: '#EFEFF5'}}
                    containerStyle={{width: 120, marginLeft: 20}}
                    raised={true}
                    onPress={() => this.props.navigation.navigate('TestPage')}
                />
                <ListItem
                    containerStyle={{height: 1000, backgroundColor: '#EFEFF5'}}
                />
            </View>
        );
    }
}
