import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './user';

@Injectable()
export class UserService {

  private url = "http://localhost:8080/api/user";
  constructor(private http: HttpClient) { }

  getUsers() {
    return this.http.get(this.url);
  }

  createUser(user: User) {
    const myHeaders = new HttpHeaders().set("Content-Type", "application/json");
    return this.http.post(this.url, JSON.stringify(user), { headers: myHeaders });
  }

  updateUser(user: User) {
    const myHeaders = new HttpHeaders().set("Content-Type", "application/json");
    return this.http.put(this.url, JSON.stringify(user), { headers: myHeaders });
  }

  deleteUser(id: string) {
    return this.http.delete(this.url + '/' + id);
  }
}
