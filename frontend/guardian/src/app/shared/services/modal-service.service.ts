import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class ModalServiceService {
    private showModalSubject: BehaviorSubject<boolean> =
        new BehaviorSubject<boolean>(false);

    showModal$: Observable<boolean> = this.showModalSubject.asObservable();

    toggleModal() {
        this.showModalSubject.next(!this.showModalSubject.value);
    }
}
