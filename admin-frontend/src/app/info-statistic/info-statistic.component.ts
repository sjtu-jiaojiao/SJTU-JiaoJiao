import { Component, OnInit } from '@angular/core';
import { EChartOption} from 'echarts';
import * as echarts from 'echarts/lib/echarts';
import { InfoService } from '../info.service';
import { buyInfo, sellInfo } from '../entity/info';
import {prepareBoxplotData} from 'echarts/extension/dataTool';
import { Transaction } from '../entity/transaction';
import { TransactionService } from '../transaction.service';
import { Format } from '../Formatter/format';
import { Router } from '@angular/router';
import { FileService } from '../file.service';
export function fdgFormatter 
    (p) {
        if(p.dataType=='node')
        return p.data.name + ' has completed '+ p.data.value + ' transaction';
        if(p.dataType=='edge')
        return p.data.source + ' has selled '+ p.data.value + ' goods to '+ p.data.target;
    }
@Component({
  selector: 'app-info-statistic',
  templateUrl: './info-statistic.component.html',
  styleUrls: ['./info-statistic.component.css']
})
export class InfoStatisticComponent implements OnInit {
  fdgoption: EChartOption;
  cldoption: any;
  tsoption: any;
  lqoption: any;
  goodoption: any;
  prcdata: any[];
  selectedInfo: buyInfo[];
  bi: buyInfo[]=[];
  si: sellInfo[]=[];
  tr: Transaction[]=[];
  pl: boolean = false;
  constructor(private router: Router, private trs: TransactionService, private is: InfoService, private fileService: FileService) { }
  ngOnInit() {
    const now = new Date().getFullYear();
    this.getAllInfo();
    this.getAllTR(new Date(now,1,1).getTime()/1000, new Date(now+1,1,1).getTime()/1000);
    this.cloudGrpah();
    this.forceGraph();
    this.calenderGraph();
    this.lineGraph();
    this.prcGraph();    
                
  }
  pauseLine(){
      this.pl=!this.pl;
  }
  getComment(){
    this.bi.forEach(
      info => {
        this.fileService.getContent(info.contentID).subscribe(
          e => {
              if(e){
              info.tags = e.tags;
              }
          }
      )
      }
    );
    if(!this.pl){
    this.prcGraph();
    this.cloudGrpah();
    }
  }
  getAllTR(beg, end){
    this.tr = this.trs.getAllTR(6,beg,end);
    if(!this.pl){
        this.calenderGraph();
       this.forceGraph();
    }
    setTimeout(() => {
        this.getAllTR(beg,end);
    }, 10000);
  }
  getAllInfo(){
    this.bi = this.is.getAllBuyInfo();
    this.bi = this.bi.sort( (a,b) => a.releaseTime - b.releaseTime);
    this.si = this.is.getAllSellInfo();
    this.si = this.si.sort( (a,b) => a.releaseTime - b.releaseTime);
    this.getComment();
    if(!this.pl)this.lineGraph();
      setTimeout(() => {
          this.getAllInfo();
      }, 10000);
  }

    onBrushSelected(param){
    this.selectedInfo = param.batch[0].selected[0].dataIndex.map(
        i => this.prcdata[i]
    );
}


  prcGraph() {
      this.prcdata=[];
    this.bi.forEach( e => 
        {
            if(e.tags && e.price)
        e.tags.forEach(
            t => {
                this.prcdata.push([t, e.price, e.buyInfoID]);
            }
        )
        }
    );
this.goodoption = {
    backgroundColor: '#01193d',
    title: [
        {
            text: '交易价格分布',
            left: 'center',
        }
    ],
    toolbox: {
			brush: {
				outOfBrush: {
					color: '#abc'
				},
				brushStyle: {
					borderWidth: 2,
					color: 'rgba(0,0,0,0.2)',
					borderColor: 'rgba(0,0,0,0.5)',
				},
				seriesIndex: [0, 1],
				throttleType: 'debounce',
				throttleDelay: 300,
				geoIndex: 0
			},
    },
	brush: {
	outOfBrush: {
		color: '#abc'
	},
	brushStyle: {
		borderWidth: 2,
		color: 'rgba(0,0,0,0.2)',
		borderColor: 'rgba(0,0,0,0.5)',
	},
	seriesIndex: [0, 1],
	throttleType: 'debounce',
	throttleDelay: 300,
	geoIndex: 'all'
	},
    tooltip: {
        trigger: 'item',
        axisPointer: {
            type: 'shadow'
        }
    },
    dataZoom: [
        {
            show: true,
            realtime: true,
            start: 0,
            end: 40
        },
        {
            type: 'inside',
            realtime: true,
            start: 0,
            end: 40
        }
    ],    
    dataset: {
        dimensions: ['tag','price','buyInfoID'],
        source: this.prcdata
    },
    grid: {
        left: '10%',
        right: '10%',
        bottom: '15%'
    },
    xAxis: {
        type: 'category',
        boundaryGap: true,
        nameGap: 30
    },
    yAxis: {
        type: 'value',
        name: 'yuan'
    },
    series: [
        {
            name: 'tag price',
            type: 'scatter',
            encode: {
                x: 'tag',
                y: 'price'
            }
        }
    ]
};
  }
  cloudGrpah() {
    const fre = new Map();
    this.bi.forEach( e => 
        {
            if(e.tags)
        e.tags.forEach(
            t => {
                if(t in fre)fre[t] +=1;
                else fre[t]=1;
            }
        )
        }
    )
    const clddata = [];
    for( const k in fre) 
    clddata.push({ name: k,
      value: fre[k],
      textStyle: {
          normal: {},
          emphasis: {}
      }
  });

    this.cldoption = {
    backgroundColor: '#01193d',
      title: {
      text: '标签热度',
      left: 'center',
  },
  tooltip: {},
  series: [{
      type: 'wordCloud',
      shape: 'circle',

      left: 'center',
      top: 'center',
      sizeRange: [12, 30],
      rotationRange: [-90, 90],
      rotationStep: 45,
      gridSize: 8,
      drawOutOfBound: false,
      textStyle: {
          normal: {
              fontFamily: 'sans-serif',
              fontWeight: 'bold',
              color: this.randamColor
          },
          emphasis: {
              shadowBlur: 10,
              shadowColor: '#333'
          }
      },
      data: clddata 
  }
  ]
};
  }
  randamColor(){return 'rgb(' + [
                      Math.round(Math.random() * 250),
                      Math.round(Math.random() * 250),
                      Math.round(Math.random() * 250)
                  ].join(',') + ')';

  }
  calenderGraph() {
    let td =new Map();
    this.tr.forEach( e => {
        if(!e)return;
       const str = Format(new Date(e.createTime*1000),'yyyy-MM-dd');
       if(str in td)
        td[str]+=1;
        else 
        td[str]=1;
    });
    let tdata = [];
    for(let i in td){
        tdata.push([i,td[i]]);
    }
    this.tsoption = {
      backgroundColor: '#01193d',
      title: {
          text: '活跃度日历',
          left: 'center'
      },
      tooltip : {
          trigger: 'item'
      },
      grid: {
        containLabel: true,
      },
      calendar: [{
          left: 80,
          range: ['2019'],
          splitLine: {
              show: true,
              lineStyle: {
                  color: '#000',
                  width: 2,
                  type: 'solid'
              }
          },
          width: '80%',
          height: '80%',
          dayLabel: {
            textStyle: {
                color: '#fff'
            }
          },
          monthLabel: {
            textStyle: {
                color: '#fff'
            }
          },
          yearLabel: {
              formatter: '{start}',
              textStyle: {
                  color: '#fff'
              }
          },
          itemStyle: {
              normal: {
                  color: '#323c48',
                  borderWidth: 1,
                  borderColor: '#111'
              }
          }
      }],
          series: [{
            type: 'effectScatter',
            coordinateSystem: 'calendar',
            symbolSize: this.cldsz,
            showEffectOn: 'render',
            rippleEffect: {
                brushType: 'stroke'
            },
            hoverAnimation: true,
            tooltip: {
                formatter: this.cldfm,
            },
            itemStyle: {
                    color: '#f4e925',
                    shadowBlur: 10,
                    shadowColor: '#333'
            },
              data: tdata
          }]
        };
  }
  cldsz(val){
                return val[1];
}
  cldfm(param) {
      return param.data[1] + ' completed transactions created in ' +param.data[0]
}
  forceGraph() {
    let td =new Map();
    let join = [];
    this.tr.forEach( e => {
        if(!e)return;
        if(e&&e.category==2){
            const other = e.toUserID;
            if(e.fromUserID in td)
                if(other in td[e.fromUserID])
                    td[e.fromUserID][other] +=1;
                else td[e.fromUserID][other] =1;
            else {
                td[e.fromUserID] = new Map();
                td[e.fromUserID][other] = 1;
            }
        }
        else if(e&& e.category==1){
        const other = e.toUserID;
        if(other in td)
            if(e.fromUserID in td[other])
                td[other][e.fromUserID] +=1;
            else td[other][e.fromUserID] =1;
        else {
            td[other] = new Map();
            td[other][e.fromUserID] = 1;
        }
    }
});
const VMAP= {};
const E = [];
const V = [];
for( let i in td){
    let sum = 0;
    for( let j in td[i]){
        E.push( 
            {source: i, target: j
            , value: td[i][j], lineStyle: {
              width: td[i][j]
            }});
        sum += td[i][j];
        if(j in VMAP)
        VMAP[j]+=td[i][j];
        else 
        VMAP[j]= td[i][j];
    }
    if(i in VMAP)
    VMAP[i]+=sum;
    else 
    VMAP[i]=sum;
}
for( let i in VMAP){
    V.push({name: i,  value : VMAP[i], symbolSize: VMAP[i]*5,
  draggable: true})
}
this.fdgoption = {
backgroundColor: '#01193d',
  title: {
      text: '交易网络',
      left: 'center',
  },
  tooltip: {
      formatter: fdgFormatter
  },
  visualMap: {
      type: 'continuous',
      min: 1,
      max: 10,
      text:['Active','Lazy'],
      realtime: false,
      calculable : true,
      color: ['orangered','yellow','lightskyblue']
  },
  animationDurationUpdate: 1500,
  animationEasingUpdate: 'quinticInOut',
  series : [
      {
          type: 'graph',
          layout: 'force',
          data: V, edges: E,
    label: {
      emphasis: {
          position: 'right',
          show: true
      }
  },
  force: {
    repulsion : 100
  },
  roam: true,
  focusNodeAdjacency: true,
  lineStyle: {
      normal: {
          color :'source',
          type : 'solid',
          width: 0.5,
          curveness: 0.2,
          opacity: 0.7
      }
  }
}
      ]
  };
  }
  fmt(t: Date) {
    return [t.getFullYear(), t.getMonth() + 1, t.getDate()].join('/');
  }
  clickForce(param){
      if(param.dataType=='node')
      this.router.navigateByUrl('/user/'+param.data.name);
  }
  lineGraph() {
    let bd =new Map();
    this.bi.forEach( e => {
        if(!e||e.releaseTime<0)return;
       const str = this.fmt(new Date(e.releaseTime*1000));
       if(str in bd)
        bd[str]+=1;
        else 
        bd[str]=1;
    });
    let sd =new Map();
    this.si.forEach( e => {
        if(!e||e.releaseTime<0)return;
       const str = this.fmt(new Date(e.releaseTime*1000));
       if(str in sd)
        sd[str]+=1;
        else 
        sd[str]=1;
    });
    let bdata = [];
    for(let i in bd){
        bdata.push([i,bd[i]]);
    }

    let sdata = [];
    for(let i in sd){
        sdata.push([i,sd[i]]);
    }
    this.lqoption = {
    backgroundColor: '#01193d',
    title : {
        text: '新增交易趋势',
        x: 'center',
        align: 'right'
    },
    grid: {
        bottom: 80
    },
    tooltip : {
        trigger: 'axis',
        axisPointer: {
            type: 'cross',
            animation: false,
            label: {
                backgroundColor: '#505765'
            }
        }
    },
    legend: {
        data: ['购买', '出售'],
        x: 'left'
    },
    dataZoom: [
        {
            show: true,
            realtime: true,
            start: 0,
            end: 100
        },
        {
            type: 'inside',
            realtime: true,
            start: 0,
            end: 100
        }
    ],
    xAxis : [
        {
            type : 'time',
            boundaryGap : false,
            axisLine: {onZero: false},
        }],
    yAxis: [
        {
            name: '新增购买信息',
            type: 'value',
            max: (value) => value.max*3/2
        },
        {
            name: '新增出售信息',
            nameLocation: 'start',
            type: 'value',
            inverse: true,
            max: (value) => value.max*3/2
        }
    ],
    series: [
        {
            name: '购买',
            yAxisIndex: 0,
            type:'line',
            smooth:true,
            symbol: 'circle',
            symbolSize: 1,
            sampling: 'average',
            itemStyle: {
                normal: {
                    color: '#d68262'
                }
            },
            areaStyle: {
                normal: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                        offset: 0,
                        color: '#d68262'
                    }, {
                        offset: 1,
                        color: '#ffe'
                    }])
                }
            },
            data: bdata
        },
        {
            name: '出售',
            type: 'line',
            yAxisIndex: 1,
            symbolSize: 5,
            smooth: true,
            symbol: 'circle',
            sampling: 'average',
            itemStyle: {
                normal: {
                    color: '#8ec6ad'
                }
            },
            areaStyle: {
                normal: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                        offset: 0,
                        color: '#8ec6ad'
                    }, {
                        offset: 1,
                        color: '#ffe'
                    }])
                }
            },
            data: sdata
        }
    ]
};

  }
}
