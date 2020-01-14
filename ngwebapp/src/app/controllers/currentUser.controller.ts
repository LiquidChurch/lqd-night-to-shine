import { Apollo } from 'apollo-angular';
import { CookieService } from 'ngx-cookie-service';
import { Component, OnInit } from '@angular/core';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, Subject } from 'rxjs';
import { map } from 'rxjs/operators';

// import { UserDetail, ItemDetail, GetUserGQL, GetTeamsGQL } from '../graphqls';
// import { RefreshUserGQL, RefreshTeamsGQL } from '../graphqls';

import { UserDetail, GetUserGQL, RefreshUserGQL } from '../graphqls';

@Injectable()
export class CurrentUserController {
  /**
   * @ignore
   */
  isDebug = true;

  /**
   * CurrentUser Object
   */
  private currentUser: Observable<UserDetail>;

  /**
   * UserTeams Object
   */
//  private userTeams: Observable<ItemDetail[]>;

  /**
   * TeamProjects Object
   */
//  private teamProjects = new Subject<ItemDetail[]>();

  /**
   * CurrentTeam Object
   */
//  private currentTeam = new Subject<ItemDetail>();

//  private initialTeamsLoad = true;
  private currentUserId: string;
//  private currentTeamId: string;

  // private isLoaded = false;

  constructor(private apollo: Apollo,
              private cookie: CookieService,
//              private getTeamsGQL: GetTeamsGQL,
              private getUserGQL: GetUserGQL,
              private refreshUserGQL: RefreshUserGQL,
//              private refreshTeamsGQL: RefreshTeamsGQL,
              private router: Router) {
  }

  async updateCurrentUser() {
    if (this.isDebug) {console.log('topbar: currentUser', this.currentUser); }
    this.currentUserId = await this.checkCurrentUser();
    if (this.currentUserId !== undefined || this.currentUserId === '') {
      console.log('1-currentUserId', this.currentUserId);
//      this.checkCurrentTeam();
    }
  }

  async checkCurrentUser(): Promise<string> {
    if (this.isDebug) { console.log('currenUserService: checkCurrentUser enter'); }
    this.currentUserLoad();
    // const uid = <string> await this.checkCurrentUserExist();

    if (this.isDebug) { console.log('currenUserService: checkCurrentUser exit'); }
    return <string> await this.checkCurrentUserExist();
  }

/* 
  async checkCurrentTeam(): Promise<boolean> {
    if (this.isDebug) { console.log('currenUserService: checkCurrentTeam enter'); }
    this.userTeamsLoad();
    const teamExist = <boolean> await this.checkCurrentTeamExist();
    if (this.isDebug) { console.log('currenUserService: checkCurrentTeam exit'); }
    return teamExist;
  }
*/
  
  /**
   * Calls [CurrentUserGQL]{@link CurrentUserGQL} and loads the CurrentUserObject
   */
  currentUserLoad(): void {
    if (this.isDebug) { console.log('currenUserService: currentUserLoad enter'); }
    this.currentUser = this.getUserGQL.watch()
        .valueChanges
        .pipe(
          map(result => result.data.getUser)
        );
    if (this.isDebug) { console.log('currenUserService: currentUserLoad exit'); }
  }

  /**
   * Refreshs current user.
   */
  currentUserRefresh(): void {
    this.refreshUserGQL
      .mutate({
      })
      .subscribe(result => {
        console.log('result', result.data.getUser.uid);
        if (result.data.getUser.uid !== this.currentUserId && result.data.getUser.uid !== '') {
//           this.checkCurrentTeam();
        }
        if (this.isDebug) { console.log('currenUserService: refreshed current user'); }
      });
  }

  /**
   * Loads User Teams
   */
/*  userTeamsLoad(): void {
    if (this.isDebug) { console.log('currenUserService: userTeamsLoad enter'); }
    if (this.initialTeamsLoad) {
      if (this.isDebug) { console.log('currenUserService: Initial Teams Load'); }
      this.userTeams = this.getTeamsGQL.watch()
          .valueChanges
          .pipe(
            map(result => result.data.getTeams)
          );
      this.initialTeamsLoad = false;
    } else {
      if (this.isDebug) { console.log('currenUserService: Refresh Teams Load'); }
      this.apollo
        .mutate({
          mutation: this.refreshTeamsGQL.document,
          variables: {},
          update: (store, { data: { getTeams }}) => {
            let data: any;
            data = store.readQuery({ query: this.getTeamsGQL.document });
            data.getTeams = [];
            for (let i = 0; i < getTeams.length; i++) {
              data.getTeams.push(getTeams[i]);
            }
            store.writeQuery({ query: this.getTeamsGQL.document, data });
          }
        })
        .subscribe();
    }
    if (this.isDebug) { console.log('currenUserService: userTeamsLoad exit'); }
  }

  checkCurrentTeamExist() {
    if (this.isDebug) { console.log('currentUserService: checkCurrentTeamExist enter'); }
    return new Promise(resolve => {
      const currentTeamId = this.cookie.get('userTeam-' + this.currentUserId);
      let hasTeam: boolean;
      if (currentTeamId === '') {
        if (this.isDebug) { console.log('currentUserService: no current team found'); }
        hasTeam = false;
      } else {
        this.userTeams.subscribe((teams) => {
          const i = teams.map(function(e) { return e.id; }).indexOf(currentTeamId);
          if (i >= 0) {
            if (this.isDebug) { console.log('currentUserService: currentTeam:', teams[i]); }
            this.currentTeam.next(teams[i]);
            this.teamProjects.next(teams[i].projects);
            console.log('teamProjects:', this.teamProjects);
            hasTeam = true;
          } else {
            if (this.isDebug) { console.log('currentUserService: no current team found'); }
            this.currentTeam.next({ id: '',
                                    description: '',
                                    pictureURL: '',
                                    type: '',
                                    role: '',
                                    name: '' });
            hasTeam = false;
          }
        });
      }
      resolve(hasTeam);
    });
    if (this.isDebug) { console.log('currentUserService: checkCurrentTeamExist exit'); }
  }
*/
  checkCurrentUserExist() {
    if (this.isDebug) { console.log('currentUserService: checkCurrentUserExist enter'); }
    return new Promise(resolve => {
      let uid: string;
      this.currentUser.subscribe((val) => {
        if (val.name !== '') {
          if (this.isDebug) { console.log('currentUserService: user found', val.uid); }
          uid = val.uid;
        } else {
          if (this.isDebug) { console.log('currentUserService: no user found'); }
          uid = '';
          this.cookie.delete('authToken', '/');
          this.cookie.delete('authType', '/');
          // this.router.navigate(['']);
        }
        if (this.currentUserId !== uid) {
          this.currentUserId = uid;
//          this.checkCurrentTeam();
        }
      });
      resolve(uid);
    });
    if (this.isDebug) { console.log('currentUserService: checkCurrentUserExist exit'); }
  }

  getCurrentUser(): Observable<UserDetail> {
    return this.currentUser;
  }
/*
  getUserTeams(): Observable<ItemDetail[]> {
    return this.userTeams;
  }

  getTeamProjects(): Observable<ItemDetail[]> {
    return this.teamProjects;
  }
  
  getCurrentTeam(): Observable<ItemDetail> {
    console.log('getCurrentTeam entered');
    console.log(this.currentTeam);
    this.currentTeam.subscribe((val) => {
      console.log('getCurrentTeam:', val.id);
    });
      console.log('getCurrentTeam exit');
    return this.currentTeam;

  }

  async setCurrentTeam(newTeam): Promise<boolean> {
    this.cookie.set('userTeam-' + this.currentUserId, newTeam, null, '/');
    const teamExist = <boolean> await this.checkCurrentTeamExist();
    return teamExist;
  }
*/
}
