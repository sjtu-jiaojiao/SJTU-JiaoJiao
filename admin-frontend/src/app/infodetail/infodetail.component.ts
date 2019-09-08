import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { sellInfo } from '../entity/info';
import { Location } from '@angular/common';
import { InfoService } from '../info.service';
import { Format } from '../Formatter/format';
import { FileService } from './../file.service';
import { Media } from '../entity/content';

@Component({
  selector: 'app-infodetail',
  templateUrl: './infodetail.component.html',
  styleUrls: ['./infodetail.component.css']
})
export class InfoDetailComponent implements OnInit {

  d = [];
  fnoption: any;
  type :string;
  state: number =0;
  contents: Media[];
  now = new Date(1997, 9, 3);
  oneDay = 24 * 3600 * 1000;
  option: any;
  deadLine: Date;
  value = Math.random() * 1000;
  @Input() info: any ;
  constructor(
  private route: ActivatedRoute,
  private infoService: InfoService,
  private location: Location,
  private fileService:FileService
) {}
stringToDate(params) {
    const date = new Date(params*1000);
    return Format(date,'yyyy-MM-dd HH:mm:ss');
    }

  ngOnInit() {
      this.graph();
      this.type = this.route.snapshot.paramMap.get('type');
      this.getinfo();
  }
    getContent(): void {
    this.fileService.getContent(this.info.contentID).subscribe(
        e => {
            if(e){
            this.contents = e.files;
            this.info.tags = e.tags;
            }
        }
    )
    }
    goBack(): void {
    this.location.back();
  }
    getinfo(): void {
    const id = parseInt(this.route.snapshot.paramMap.get('id'));
    if(this.type === 'sellInfo')
    this.infoService.getSellInfo(id)
      .subscribe(info => {
          if(!info)return;
          this.info = info;
          this.deadLine = new Date(this.info.validTime*1000);
          this.getContent();
      });
    else if(this.type === 'buyInfo')
    this.infoService.getBuyInfo(id)
      .subscribe(info => {
        if(!info)return;
          this.info = info;
          this.deadLine = new Date(this.info.validTime*1000);
          this.getContent();
      });
  }

    save(): void {
    if(!this.info) return;
    this.info.validTime = this.deadLine.getTime()/1000;
    if(this.type==='sellInfo')
    this.infoService.updateSellInfo(this.info)
        .subscribe(() => this.goBack());
    else if(this.type ==='buyInfo')
    this.infoService.updateBuyInfo(this.info)
        .subscribe(() => this.goBack());
    }

  graph() {
    this.fnoption = {
    title: {
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

}
}
