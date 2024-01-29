import { Component, OnDestroy } from '@angular/core';
import { Observable, Subscription, combineLatest, map } from 'rxjs';
import { Prescription } from 'src/app/core/models/rx/Prescription';
import { ModalServiceService } from 'src/app/shared/services/modal-service.service';
import { PrescriptionService } from 'src/app/shared/services/prescription/prescription.service';

@Component({
    selector: 'app-rx-edit-modal',
    templateUrl: './rx-edit-modal.component.html',
    styleUrls: ['./rx-edit-modal.component.css'],
})
export class RxEditModalComponent implements OnDestroy {
    showModal$: Observable<boolean>;
    rx$: Observable<Prescription>;
    data$: Observable<{ showModal: boolean; rx: Prescription }>;

    private modalSubscription: Subscription;
    private rxSubscription: Subscription;

    constructor(
        private modalService: ModalServiceService,
        private rxService: PrescriptionService
    ) {
        this.showModal$ = this.modalService.showModal$;
        this.modalSubscription = this.showModal$.subscribe();

        this.rx$ = this.rxService.showPrescription$;
        this.rxSubscription = this.rx$.subscribe();

        this.data$ = combineLatest([this.showModal$, this.rx$]).pipe(
            map(([showModal, rx]) => ({
                showModal,
                rx,
            }))
        );
    }

    toggleModal() {
        this.modalService.toggleModal();
    }

    ngOnDestroy() {
        this.modalSubscription.unsubscribe();
        this.rxSubscription.unsubscribe();
    }
}
