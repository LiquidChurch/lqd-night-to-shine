import { Component, OnInit, AfterViewInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router, NavigationEnd, ActivatedRoute } from '@angular/router';
import { filter, map } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {
  appTitle = 'GAE-SPA';
  clientHeight: number;
  clientWidth: number;
  footerHeight: number;
  
  constructor(private title: Title,
              private router: Router,
              private activatedRoute: ActivatedRoute) {
    this.clientHeight = window.innerHeight;
    this.clientWidth = window.innerWidth;
  }
  
  ngOnInit() {
    const appTitle = this.title.getTitle();
    
    this.router
      .events.pipe(
      filter(event => event instanceof NavigationEnd),
      map(() => {
        const child = this.activatedRoute.firstChild;
        if (child.snapshot.data['title']) {
          return child.snapshot.data['title'];
        }
        return appTitle;
      })
     ).subscribe((ttl: string) => {
      this.title.setTitle(ttl + " - " + this.appTitle);
    });
  }
  
  ngAfterViewInit() {
    console.log('AppComponent after view init');
    if (this.clientWidth < 640 ) {
      // this.footerHeight = document.getElementById('bottombar').offsetHeight - 3;
      this.footerHeight = 63.39;      
      this.clientHeight = this.clientHeight - this.footerHeight;
    } else {
      this.footerHeight = 63.39;
      this.clientHeight = this.clientHeight - this.footerHeight;
    }
  }  
  
  setPageTitle(pageTitle: string) {
    this.title.setTitle(pageTitle);
  }
}