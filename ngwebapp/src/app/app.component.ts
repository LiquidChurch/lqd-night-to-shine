import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router, NavigationEnd, ActivatedRoute } from '@angular/router';
import { filter, map } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {
  defaultTitle = 'GAE SPA Base';
  
  constructor(private title: Title,
              private router: Router,
              private activatedRoute: ActivatedRoute) {}
  
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
      this.title.setTitle(ttl + " - " + this.defaultTitle);
    });
  }
  
  setPageTitle(pageTitle: string) {
    this.title.setTitle(pageTitle);
  }
}