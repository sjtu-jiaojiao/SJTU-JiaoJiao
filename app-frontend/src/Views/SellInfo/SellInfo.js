import React, { Component } from 'react';
import {FlatList, StyleSheet, Text, TouchableOpacity, View} from 'react-native';
import Config from "../../Config";
import {ListItem} from "react-native-elements";
import {parseStatus, parseTimeStamp} from "../../Utils/ParseInfo";

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

export default class SellInfoScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            SellInfoList: null,
            loaded: false,
        };
        this.fetchData = this.fetchData.bind(this);
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        //console.warn(obj);
        Http.get('/sellInfo', obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    SellInfoList: response.sellInfo,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>出售信息</Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => {
        let sellInfoID = item.sellInfoID;
        let header = ('出售：' + item.goodName);
        let infoType = 'sellInfo';
        return (
            <TouchableOpacity onPress={() => this.props.navigation.push('GoodInfo', { infoType, sellInfoID, header })}>
                <ListItem
                    bottomDivider
                    containerStyle={{height: 200}}
                    title={
                        <Text numberOfLines={1} style={{color: 'black', fontWeight: 'bold', fontSize: 20}}>出售物品：{item.goodName}</Text>
                    }
                    subtitle={
                        <View style={styles.subtitleView}>
                            <Text numberOfLines={1} style={styles.ratingText}>具体描述：{item.description}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>信息状态：{parseStatus(item.status)}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>出售价格：￥{item.price}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>发布时间：{parseTimeStamp(item.releaseTime)}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>物品标签：暂无</Text>
                        </View>
                    }
                />
            </TouchableOpacity>
        );
    };

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else {
            //console.warn(this.state.SellInfoList[0].releaseTime);
            //console.warn(TimeStamptoDate(this.state.SellInfoList[0].releaseTime));
            //console.warn(TimeStampNow());
            //console.warn(TimeStamptoDate(TimeStampNow()));
            return (
                <View>
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.SellInfoList}
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

