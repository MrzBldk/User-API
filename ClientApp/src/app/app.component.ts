import { Component, OnInit, ViewChild, TemplateRef } from '@angular/core';
import { UserService } from './user.service';
import { User } from './user';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  providers: [UserService]
})
export class AppComponent implements OnInit {
  @ViewChild('readOnlyTemplate', { static: false }) readOnlyTemplate: TemplateRef<any> | null = null
  @ViewChild('editTemplate', { static: false }) editTemplate: TemplateRef<any> | null = null

  editedUser: User | null = null;
  users: Array<User>;
  isNewRecord: boolean = false;
  title = 'ClientApp'

  constructor(private serv: UserService) {
    this.users = new Array<User>();
  }

  ngOnInit() {
    this.loadUsers();
  }

  private loadUsers() {
    this.serv.getUsers().subscribe((data: any) => {
      this.users = data.data;
    });
  }

  addUser() {
    this.editedUser = new User("", "", 0);
    this.users.push(this.editedUser);
    this.isNewRecord = true;
  }

  editUser(user: User) {
    this.editedUser = new User(user.id, user.name, user.age);
  }

  loadTemplate(user: User) {
    if (this.editedUser && this.editedUser.id === user.id) {
      return this.editTemplate;
    } else {
      return this.readOnlyTemplate;
    }
  }

  saveUser() {
    if (this.isNewRecord) {
      this.serv.createUser(this.editedUser as User).subscribe((data: any) => {
        this.loadUsers();
      });
      this.isNewRecord = false;
    } else {
      this.serv.updateUser(this.editedUser as User).subscribe((data: any) => {
        this.loadUsers();
      });
    }
    this.editedUser = null;
  }

  cancel() {
    if (this.isNewRecord) {
      this.users.pop();
      this.isNewRecord = false;
    }
    this.editedUser = null;
  }

  deleteUser(user: User) {
    this.serv.deleteUser(user.id).subscribe((data: any) => {
      this.loadUsers();
    });
  }
}
