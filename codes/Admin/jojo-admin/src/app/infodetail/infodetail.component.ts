import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-infodetail',
  templateUrl: './infodetail.component.html',
  styleUrls: ['./infodetail.component.css']
})
export class InfoDetailComponent implements OnInit {

  d = [];
<<<<<<< HEAD
  fnoption: any;
=======
>>>>>>> cbf88643b148d3696f6eaf12c3d034083aee3317
  now = new Date(1997, 9, 3);
  oneDay = 24 * 3600 * 1000;
  option: any;
  value = Math.random() * 1000;
  randomData() {
    this.now = new Date(+this.now + this.oneDay);
    this.value = this.value + Math.random() * 21 - 10;
    return {
        name: this.now.toString(),
        value: [
            [this.now.getFullYear(), this.now.getMonth() + 1, this.now.getDate()].join('/'),
            Math.round(this.value)
        ]
    };
}

  constructor() { }

  ngOnInit() {
    for (let i = 0; i < 1000; i++) {
        this.d.push(this.randomData());
}
    this.option = {
    title: {
        text: '价格波动曲线'
    },
    tooltip: {
        trigger: 'axis',
        formatter: (params) => {
            params = params[0];
            const date = new Date(params.name);
            return date.getDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear() + ' : ' + params.value[1];
        },
        axisPointer: {
            animation: false
        }
    },
    xAxis: {
        type: 'time',
        splitLine: {
            show: false
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
            show: false
        }
    },
    series: [{
        name: '模拟数据',
        type: 'line',
        showSymbol: false,
        hoverAnimation: false,
        data: this.d
    }]
};

    setInterval(() => {

    for (let i = 0; i < 5; i++) {
        this.d.shift();
        this.d.push(this.randomData());
    }
    this.fnoption = {
        title: {
            text: '流量跟踪漏斗图',
        },
        tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b} : {c}'
        },
        legend: {
            data: ['展现', '点击' , '申请' , '预约' , '完成']
        },
        calculable: true,
        series: [
            {
                name: '漏斗图',
                type: 'funnel',
                left: '10%',
                top: 60,
                bottom: 60,
                width: '80%',
                min: 0,
                max: 100,
                minSize: '0%',
                maxSize: '100%',
                sort: 'descending',
                gap: 2,
                label: {
                    show: true,
                    position: 'inside'
                },
                labelLine: {
                    length: 10,
                    lineStyle: {
                        width: 1,
                        type: 'solid'
                    }
                },
                itemStyle: {
                    borderColor: '#fff',
                    borderWidth: 1
                },
                emphasis: {
                    label: {
                        fontSize: 20
                    }
                },
                data: [
                    {value: 1, name: '完成'},
                    {value: 2, name: '预约'},
                    {value: 10, name: '申请'},
                    {value: 100, name: '点击'},
                    {value: 1000, name: '展现'}
                ]
            }
        ]
    };
    this.option = {
      title: {
          text: '价格波动曲线'
      },
      tooltip: {
          trigger: 'axis',
          formatter: (params) => {
              params = params[0];
              const date = new Date(params.name);
              return date.getDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear() + ' : ' + params.value[1];
          },
          axisPointer: {
              animation: false
          }
      },
      xAxis: {
          type: 'time',
          splitLine: {
              show: false
          }
      },
      yAxis: {
          type: 'value',
          boundaryGap: [0, '100%'],
          splitLine: {
              show: false
          }
      },
      series: [{
          name: '模拟数据',
          type: 'line',
          showSymbol: false,
          hoverAnimation: false,
          data: this.d
      }]
  };
}, 1000);
  }

}
