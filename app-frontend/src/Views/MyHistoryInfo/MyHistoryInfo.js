import React, { Component } from 'react';
import {FlatList, Text, View, StyleSheet} from 'react-native';
import {Avatar, Button, ListItem, SearchBar, Icon } from "react-native-elements";
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

export default class MyHistoryInfoScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            HistoryInfoList: [
                {
                    description: "求购线性代数",
                    infoType: 'Sell',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "20",
                    dealTime: '2019.03.21',
                    tradingPerson: '王新哲'
                },
                {
                    description: "求购线购线购线购线购线购线性代数",
                    infoType: 'Buy',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "2019.03.05",
                    tradingPerson: '康欢'
                },
                {
                    description: "求购数购数购数购数购数",
                    infoType: 'Sell',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "2019.03.05",
                    tradingPerson: '亢虎权'
                },
                {
                    description: "线性代数有没有",
                    infoType: 'Sell',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "2019.03.05",
                    tradingPerson: '朱文杰'
                },
                {
                    description: "概率统计",
                    infoType: 'Sell',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "2019.03.05",
                    tradingPerson: '林江浩'
                },
                {
                    description: "数学分析",
                    infoType: 'Buy',
                    HistoryInfoID: 1,
                    price: 12.02,
                    releaseTime: "2019.02.28",
                    validTime: "2019.03.05",
                    tradingPerson: '王新哲'
                },
            ],
            loaded: true,
        };
        //this.fetchData = this.fetchData.bind(this);
    };

    componentDidMount() {
        //this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = { userID: Config.userInfo.userID };
        //console.warn(obj);
        Http.get('/HistoryInfo', obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    HistoryInfoList: response.historyInfo,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>我的历史交易信息</Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => {
        if (item.infoType === 'Sell') {
            return (
                <ListItem
                    bottomDivider
                    containerStyle={{height: 200}}
                    title={
                        <Text numberOfLines={1} style={{color: 'black', fontWeight: 'bold', fontSize: 20}}>出售：{item.goodName}</Text>
                    }
                    subtitle={
                        <View style={styles.subtitleView}>
                            <Text numberOfLines={1} style={styles.ratingText}>出售描述：￥{item.description}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>成交价格：￥{item.price}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>发布时间：{item.releaseTime}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>成交时间：{item.validTime}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>交易对象：{item.tradingPerson}</Text>
                        </View>
                    }
                />
            );
        } else {
            return (
                <ListItem
                    bottomDivider
                    containerStyle={{height: 200}}
                    title={
                        <Text numberOfLines={1} style={{color: 'black', fontWeight: 'bold', fontSize: 20}}>求购：{item.goodName}</Text>
                    }
                    subtitle={
                        <View style={styles.subtitleView}>
                            <Text numberOfLines={1} style={styles.ratingText}>求购描述：￥{item.description}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>成交价格：￥{item.price}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>发布时间：{item.releaseTime}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>成交时间：{item.validTime}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>交易对象：{item.tradingPerson}</Text>
                        </View>
                    }
                />
            );
        }
    }

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else if (this.state.HistoryInfoList == undefined) {
            return (
                <View style={styles.container}>
                    <Text>您暂时没有发布任何历史交易信息</Text>
                </View>
            )
        }
        else {
            return (
                <View>
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.HistoryInfoList}
                        renderItem={this.renderItem}
                    />
                </View>
            )
        }
    }
}

let styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#EFEFF5"
    },
    subtitleView: {
        //flexDirection: 'row',
        paddingLeft: 10,
        paddingTop: 5
    },
    ratingText: {
        paddingLeft: 10,
        color: 'grey',
        fontSize: 16,
    },
});
