import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-infodetail',
  templateUrl: './infodetail.component.html',
  styleUrls: ['./infodetail.component.css']
})
export class InfoDetailComponent implements OnInit {

  d = [];
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
